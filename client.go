package go_shopware_admin_sdk

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/oauth2"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

var errNonNilContext = errors.New("Context must be non-nil")

type Client struct {
	url    string
	client *http.Client

	common           ClientService
	Repository       Repository
	Bulk             *BulkService
	Info             *InfoService
	ExtensionManager *ExtensionManagerService
}

type ClientService struct {
	Client *Client
}

func NewApiClient(ctx context.Context, shopUrl string, credentials OAuthCredentials, httpClient *http.Client) (*Client, error) {
	shopClient := &Client{url: shopUrl, client: httpClient}
	shopClient.common.Client = shopClient

	shopClient.Repository = NewRepository(shopClient.common)
	shopClient.Bulk = (*BulkService)(&shopClient.common)
	shopClient.Info = (*InfoService)(&shopClient.common)
	shopClient.ExtensionManager = (*ExtensionManagerService)(&shopClient.common)

	if err := shopClient.authorize(ctx, shopUrl, credentials); err != nil {
		return nil, err
	}

	return shopClient, nil
}

func (c *Client) authorize(ctx context.Context, url string, credentials OAuthCredentials) error {
	if c.client != nil {
		ctx = context.WithValue(ctx, oauth2.HTTPClient, c.client)
	}

	tokenSrc, err := credentials.GetTokenSource(ctx, url+"/api/oauth/token")
	if err != nil {
		return err
	}
	c.client = oauth2.NewClient(ctx, tokenSrc)
	return nil
}

func (c *Client) newRequest(ctx context.Context, method, path string, body io.Reader) (*http.Request, error) {
	return http.NewRequestWithContext(ctx, method, c.url+path, body)
}

func (c *Client) BareDo(ctx context.Context, req *http.Request) (*http.Response, error) {
	if ctx == nil {
		return nil, errNonNilContext
	}

	resp, err := c.client.Do(req)

	if err != nil {
		// If we got an error, and the Context has been canceled,
		// the Context's error is probably more useful.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
		return nil, err
	}

	err = checkResponse(resp)

	return resp, err
}

func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.BareDo(ctx, req)
	if err != nil {
		return resp, err
	}
	defer resp.Body.Close()

	switch v := v.(type) {
	case nil:
	case io.Writer:
		_, err = io.Copy(v, resp.Body)
	default:
		decErr := json.NewDecoder(resp.Body).Decode(v)
		if decErr == io.EOF {
			decErr = nil // ignore EOF errors caused by empty response body
		}
		if decErr != nil {
			err = decErr
		}
	}
	return resp, err
}

func (c *Client) NewRequest(context ApiContext, method, urlStr string, body interface{}) (*http.Request, error) {
	if strings.HasSuffix(c.url, "/") {
		return nil, fmt.Errorf("BaseURL must not have a trailing slash, but %q does not", c.url)
	}

	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequestWithContext(context.Context, method, c.url+urlStr, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("sw-language-id", context.LanguageId)
	req.Header.Set("sw-version-id", context.VersionId)

	if context.SkipFlows {
		req.Header.Set("sw-skip-trigger-flow", "1")
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("Accept", "application/json")

	return req, nil
}

func checkResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}

	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		_ = json.Unmarshal(data, errorResponse)
	}

	errorResponse.Content = string(data)

	return errorResponse
}

type ErrorResponse struct {
	Response *http.Response
	Content  string
	Errors   []ErrorDetail `json:"errors"`
}

type ErrorDetail struct {
	Code   string `json:"code"`
	Status string `json:"status"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

func (r ErrorResponse) Error() string {
	detail := r.Content

	for _, detailError := range r.Errors {
		detail = detailError.Detail
	}

	return fmt.Sprintf("API request failed, got http code %d with content: %s", r.Response.StatusCode, detail)
}

type ApiContext struct {
	Context    context.Context
	LanguageId string
	VersionId  string
	SkipFlows  bool
}

func NewApiContext(ctx context.Context) ApiContext {
	return ApiContext{
		Context:    ctx,
		LanguageId: "2fbb5fe2e29a4d70aa5854ce7ce3e20b",
		VersionId:  "0fa91ce3e96a4bc2be4bd9ce752c3425",
		SkipFlows:  false,
	}
}

type EntityCollection struct {
	Total        int64       `json:"total"`
	Aggregations interface{} `json:"aggregations"`
}

type SearchIdsResponse struct {
	Total int      `json:"total"`
	Data  []string `json:"data"`
}

func (s SearchIdsResponse) FirstId() string {
	if len(s.Data) > 0 {
		return s.Data[0]
	}

	return ""
}

type entityDelete struct {
	Id string `json:"id"`
}

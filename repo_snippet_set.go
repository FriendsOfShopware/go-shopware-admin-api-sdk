package go_shopware_admin_sdk

import (
	"net/http"

	"time"
)

type SnippetSetRepository ClientService

func (t SnippetSetRepository) Search(ctx ApiContext, criteria Criteria) (*SnippetSetCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/snippet-set", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(SnippetSetCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t SnippetSetRepository) SearchAll(ctx ApiContext, criteria Criteria) (*SnippetSetCollection, *http.Response, error) {
	if criteria.Limit == 0 {
		criteria.Limit = 50
	}

	if criteria.Page == 0 {
		criteria.Page = 1
	}

	c, resp, err := t.Search(ctx, criteria)

	if err != nil {
		return c, resp, err
	}

	for {
		criteria.Page++

		nextC, nextResp, nextErr := t.Search(ctx, criteria)

		if nextErr != nil {
			return c, nextResp, nextErr
		}

		if len(nextC.Data) == 0 {
			break
		}

		c.Data = append(c.Data, nextC.Data...)
	}

	c.Total = int64(len(c.Data))

	return c, resp, err
}

func (t SnippetSetRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/snippet-set", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(SearchIdsResponse)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t SnippetSetRepository) Upsert(ctx ApiContext, entity []SnippetSet) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"snippet_set": {
		Entity:  "snippet_set",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t SnippetSetRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"snippet_set": {
		Entity:  "snippet_set",
		Action:  "delete",
		Payload: payload,
	}})
}

type SnippetSet struct {
	BaseFile string `json:"baseFile,omitempty"`

	Iso string `json:"iso,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	Snippets []Snippet `json:"snippets,omitempty"`

	SalesChannelDomains []SalesChannelDomain `json:"salesChannelDomains,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type SnippetSetCollection struct {
	EntityCollection

	Data []SnippetSet `json:"data"`
}

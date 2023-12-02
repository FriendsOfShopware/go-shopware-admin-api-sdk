package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type ProductDownloadRepository ClientService

func (t ProductDownloadRepository) Search(ctx ApiContext, criteria Criteria) (*ProductDownloadCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/product-download", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(ProductDownloadCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ProductDownloadRepository) SearchAll(ctx ApiContext, criteria Criteria) (*ProductDownloadCollection, *http.Response, error) {
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

func (t ProductDownloadRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/product-download", criteria)

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

func (t ProductDownloadRepository) Upsert(ctx ApiContext, entity []ProductDownload) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_download": {
		Entity:  "product_download",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t ProductDownloadRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_download": {
		Entity:  "product_download",
		Action:  "delete",
		Payload: payload,
	}})
}

type ProductDownload struct {

	VersionId      string  `json:"versionId,omitempty"`

	ProductId      string  `json:"productId,omitempty"`

	MediaId      string  `json:"mediaId,omitempty"`

	Product      *Product  `json:"product,omitempty"`

	Media      *Media  `json:"media,omitempty"`

	Id      string  `json:"id,omitempty"`

	Position      float64  `json:"position,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	ProductVersionId      string  `json:"productVersionId,omitempty"`

}

type ProductDownloadCollection struct {
	EntityCollection

	Data []ProductDownload `json:"data"`
}

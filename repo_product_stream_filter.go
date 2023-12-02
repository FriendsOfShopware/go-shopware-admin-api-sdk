package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type ProductStreamFilterRepository ClientService

func (t ProductStreamFilterRepository) Search(ctx ApiContext, criteria Criteria) (*ProductStreamFilterCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/product-stream-filter", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(ProductStreamFilterCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ProductStreamFilterRepository) SearchAll(ctx ApiContext, criteria Criteria) (*ProductStreamFilterCollection, *http.Response, error) {
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

func (t ProductStreamFilterRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/product-stream-filter", criteria)

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

func (t ProductStreamFilterRepository) Upsert(ctx ApiContext, entity []ProductStreamFilter) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_stream_filter": {
		Entity:  "product_stream_filter",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t ProductStreamFilterRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_stream_filter": {
		Entity:  "product_stream_filter",
		Action:  "delete",
		Payload: payload,
	}})
}

type ProductStreamFilter struct {

	Queries      []ProductStreamFilter  `json:"queries,omitempty"`

	Type      string  `json:"type,omitempty"`

	Field      string  `json:"field,omitempty"`

	Parameters      interface{}  `json:"parameters,omitempty"`

	Value      string  `json:"value,omitempty"`

	Position      float64  `json:"position,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Id      string  `json:"id,omitempty"`

	ProductStreamId      string  `json:"productStreamId,omitempty"`

	Operator      string  `json:"operator,omitempty"`

	ParentId      string  `json:"parentId,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	ProductStream      *ProductStream  `json:"productStream,omitempty"`

	Parent      *ProductStreamFilter  `json:"parent,omitempty"`

}

type ProductStreamFilterCollection struct {
	EntityCollection

	Data []ProductStreamFilter `json:"data"`
}

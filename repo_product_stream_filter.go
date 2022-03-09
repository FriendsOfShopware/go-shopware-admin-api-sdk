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
	Operator string `json:"operator,omitempty"`

	Parameters interface{} `json:"parameters,omitempty"`

	ProductStream *ProductStream `json:"productStream,omitempty"`

	Id string `json:"id,omitempty"`

	Queries []ProductStreamFilter `json:"queries,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Position float64 `json:"position,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	ProductStreamId string `json:"productStreamId,omitempty"`

	ParentId string `json:"parentId,omitempty"`

	Field string `json:"field,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	Type string `json:"type,omitempty"`

	Value string `json:"value,omitempty"`

	Parent *ProductStreamFilter `json:"parent,omitempty"`
}

type ProductStreamFilterCollection struct {
	EntityCollection

	Data []ProductStreamFilter `json:"data"`
}

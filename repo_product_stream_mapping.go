package go_shopware_admin_sdk

import (
	"net/http"
)

type ProductStreamMappingRepository ClientService

func (t ProductStreamMappingRepository) Search(ctx ApiContext, criteria Criteria) (*ProductStreamMappingCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/product-stream-mapping", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(ProductStreamMappingCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ProductStreamMappingRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/product-stream-mapping", criteria)

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

func (t ProductStreamMappingRepository) Upsert(ctx ApiContext, entity []ProductStreamMapping) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_stream_mapping": {
		Entity:  "product_stream_mapping",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t ProductStreamMappingRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_stream_mapping": {
		Entity:  "product_stream_mapping",
		Action:  "delete",
		Payload: payload,
	}})
}

type ProductStreamMapping struct {
	ProductId string `json:"productId,omitempty"`

	ProductVersionId string `json:"productVersionId,omitempty"`

	ProductStreamId string `json:"productStreamId,omitempty"`

	Product *Product `json:"product,omitempty"`

	ProductStream *ProductStream `json:"productStream,omitempty"`
}

type ProductStreamMappingCollection struct {
	EntityCollection

	Data []ProductStreamMapping `json:"data"`
}

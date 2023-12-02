package go_shopware_admin_sdk

import (
	"net/http"
	"time"
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

func (t ProductStreamMappingRepository) SearchAll(ctx ApiContext, criteria Criteria) (*ProductStreamMappingCollection, *http.Response, error) {
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

	ProductStreamId      string  `json:"productStreamId,omitempty"`

	Product      *Product  `json:"product,omitempty"`

	ProductStream      *ProductStream  `json:"productStream,omitempty"`

	ProductId      string  `json:"productId,omitempty"`

	ProductVersionId      string  `json:"productVersionId,omitempty"`

}

type ProductStreamMappingCollection struct {
	EntityCollection

	Data []ProductStreamMapping `json:"data"`
}

package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type ProductTagRepository ClientService

func (t ProductTagRepository) Search(ctx ApiContext, criteria Criteria) (*ProductTagCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/product-tag", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(ProductTagCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ProductTagRepository) SearchAll(ctx ApiContext, criteria Criteria) (*ProductTagCollection, *http.Response, error) {
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

func (t ProductTagRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/product-tag", criteria)

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

func (t ProductTagRepository) Upsert(ctx ApiContext, entity []ProductTag) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_tag": {
		Entity:  "product_tag",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t ProductTagRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_tag": {
		Entity:  "product_tag",
		Action:  "delete",
		Payload: payload,
	}})
}

type ProductTag struct {

	ProductVersionId      string  `json:"productVersionId,omitempty"`

	TagId      string  `json:"tagId,omitempty"`

	Product      *Product  `json:"product,omitempty"`

	Tag      *Tag  `json:"tag,omitempty"`

	ProductId      string  `json:"productId,omitempty"`

}

type ProductTagCollection struct {
	EntityCollection

	Data []ProductTag `json:"data"`
}

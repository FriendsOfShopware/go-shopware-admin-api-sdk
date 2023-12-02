package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type ProductOptionRepository ClientService

func (t ProductOptionRepository) Search(ctx ApiContext, criteria Criteria) (*ProductOptionCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/product-option", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(ProductOptionCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ProductOptionRepository) SearchAll(ctx ApiContext, criteria Criteria) (*ProductOptionCollection, *http.Response, error) {
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

func (t ProductOptionRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/product-option", criteria)

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

func (t ProductOptionRepository) Upsert(ctx ApiContext, entity []ProductOption) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_option": {
		Entity:  "product_option",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t ProductOptionRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_option": {
		Entity:  "product_option",
		Action:  "delete",
		Payload: payload,
	}})
}

type ProductOption struct {

	ProductVersionId      string  `json:"productVersionId,omitempty"`

	OptionId      string  `json:"optionId,omitempty"`

	Product      *Product  `json:"product,omitempty"`

	Option      *PropertyGroupOption  `json:"option,omitempty"`

	ProductId      string  `json:"productId,omitempty"`

}

type ProductOptionCollection struct {
	EntityCollection

	Data []ProductOption `json:"data"`
}

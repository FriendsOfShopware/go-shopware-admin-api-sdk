package go_shopware_admin_sdk

import (
	"net/http"
)

type ProductPropertyRepository ClientService

func (t ProductPropertyRepository) Search(ctx ApiContext, criteria Criteria) (*ProductPropertyCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/product-property", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(ProductPropertyCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ProductPropertyRepository) SearchAll(ctx ApiContext, criteria Criteria) (*ProductPropertyCollection, *http.Response, error) {
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

func (t ProductPropertyRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/product-property", criteria)

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

func (t ProductPropertyRepository) Upsert(ctx ApiContext, entity []ProductProperty) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_property": {
		Entity:  "product_property",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t ProductPropertyRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_property": {
		Entity:  "product_property",
		Action:  "delete",
		Payload: payload,
	}})
}

type ProductProperty struct {
	Product *Product `json:"product,omitempty"`

	Option *PropertyGroupOption `json:"option,omitempty"`

	ProductId string `json:"productId,omitempty"`

	ProductVersionId string `json:"productVersionId,omitempty"`

	OptionId string `json:"optionId,omitempty"`
}

type ProductPropertyCollection struct {
	EntityCollection

	Data []ProductProperty `json:"data"`
}

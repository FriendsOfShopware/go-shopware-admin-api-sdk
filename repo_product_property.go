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
	ProductId string `json:"productId,omitempty"`

	ProductVersionId string `json:"productVersionId,omitempty"`

	OptionId string `json:"optionId,omitempty"`

	Product *Product `json:"product,omitempty"`

	Option *PropertyGroupOption `json:"option,omitempty"`
}

type ProductPropertyCollection struct {
	EntityCollection

	Data []ProductProperty `json:"data"`
}

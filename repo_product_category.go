package go_shopware_admin_sdk

import (
	"net/http"
)

type ProductCategoryRepository ClientService

func (t ProductCategoryRepository) Search(ctx ApiContext, criteria Criteria) (*ProductCategoryCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/product-category", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(ProductCategoryCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ProductCategoryRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/product-category", criteria)

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

func (t ProductCategoryRepository) Upsert(ctx ApiContext, entity []ProductCategory) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_category": {
		Entity:  "product_category",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t ProductCategoryRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_category": {
		Entity:  "product_category",
		Action:  "delete",
		Payload: payload,
	}})
}

type ProductCategory struct {
	ProductId string `json:"productId,omitempty"`

	ProductVersionId string `json:"productVersionId,omitempty"`

	CategoryId string `json:"categoryId,omitempty"`

	CategoryVersionId string `json:"categoryVersionId,omitempty"`

	Product *Product `json:"product,omitempty"`

	Category *Category `json:"category,omitempty"`
}

type ProductCategoryCollection struct {
	EntityCollection

	Data []ProductCategory `json:"data"`
}

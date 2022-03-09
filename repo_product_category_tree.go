package go_shopware_admin_sdk

import (
	"net/http"
)

type ProductCategoryTreeRepository ClientService

func (t ProductCategoryTreeRepository) Search(ctx ApiContext, criteria Criteria) (*ProductCategoryTreeCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/product-category-tree", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(ProductCategoryTreeCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ProductCategoryTreeRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/product-category-tree", criteria)

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

func (t ProductCategoryTreeRepository) Upsert(ctx ApiContext, entity []ProductCategoryTree) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_category_tree": {
		Entity:  "product_category_tree",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t ProductCategoryTreeRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_category_tree": {
		Entity:  "product_category_tree",
		Action:  "delete",
		Payload: payload,
	}})
}

type ProductCategoryTree struct {
	CategoryVersionId string `json:"categoryVersionId,omitempty"`

	Product *Product `json:"product,omitempty"`

	Category *Category `json:"category,omitempty"`

	ProductId string `json:"productId,omitempty"`

	ProductVersionId string `json:"productVersionId,omitempty"`

	CategoryId string `json:"categoryId,omitempty"`
}

type ProductCategoryTreeCollection struct {
	EntityCollection

	Data []ProductCategoryTree `json:"data"`
}

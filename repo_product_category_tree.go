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

func (t ProductCategoryTreeRepository) SearchAll(ctx ApiContext, criteria Criteria) (*ProductCategoryTreeCollection, *http.Response, error) {
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
	ProductId string `json:"productId,omitempty"`

	ProductVersionId string `json:"productVersionId,omitempty"`

	CategoryId string `json:"categoryId,omitempty"`

	CategoryVersionId string `json:"categoryVersionId,omitempty"`

	Product *Product `json:"product,omitempty"`

	Category *Category `json:"category,omitempty"`
}

type ProductCategoryTreeCollection struct {
	EntityCollection

	Data []ProductCategoryTree `json:"data"`
}

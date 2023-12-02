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

func (t ProductCategoryRepository) SearchAll(ctx ApiContext, criteria Criteria) (*ProductCategoryCollection, *http.Response, error) {
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
	Category *Category `json:"category,omitempty"`

	ProductId string `json:"productId,omitempty"`

	ProductVersionId string `json:"productVersionId,omitempty"`

	CategoryId string `json:"categoryId,omitempty"`

	CategoryVersionId string `json:"categoryVersionId,omitempty"`

	Product *Product `json:"product,omitempty"`
}

type ProductCategoryCollection struct {
	EntityCollection

	Data []ProductCategory `json:"data"`
}

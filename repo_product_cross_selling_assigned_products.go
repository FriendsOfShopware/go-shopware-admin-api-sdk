package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type ProductCrossSellingAssignedProductsRepository ClientService

func (t ProductCrossSellingAssignedProductsRepository) Search(ctx ApiContext, criteria Criteria) (*ProductCrossSellingAssignedProductsCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/product-cross-selling-assigned-products", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(ProductCrossSellingAssignedProductsCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ProductCrossSellingAssignedProductsRepository) SearchAll(ctx ApiContext, criteria Criteria) (*ProductCrossSellingAssignedProductsCollection, *http.Response, error) {
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

func (t ProductCrossSellingAssignedProductsRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/product-cross-selling-assigned-products", criteria)

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

func (t ProductCrossSellingAssignedProductsRepository) Upsert(ctx ApiContext, entity []ProductCrossSellingAssignedProducts) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_cross_selling_assigned_products": {
		Entity:  "product_cross_selling_assigned_products",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t ProductCrossSellingAssignedProductsRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_cross_selling_assigned_products": {
		Entity:  "product_cross_selling_assigned_products",
		Action:  "delete",
		Payload: payload,
	}})
}

type ProductCrossSellingAssignedProducts struct {

	ProductId      string  `json:"productId,omitempty"`

	Product      *Product  `json:"product,omitempty"`

	CrossSelling      *ProductCrossSelling  `json:"crossSelling,omitempty"`

	Position      float64  `json:"position,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Id      string  `json:"id,omitempty"`

	ProductVersionId      string  `json:"productVersionId,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CrossSellingId      string  `json:"crossSellingId,omitempty"`

}

type ProductCrossSellingAssignedProductsCollection struct {
	EntityCollection

	Data []ProductCrossSellingAssignedProducts `json:"data"`
}

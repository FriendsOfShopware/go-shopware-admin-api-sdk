package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type ProductCrossSellingRepository ClientService

func (t ProductCrossSellingRepository) Search(ctx ApiContext, criteria Criteria) (*ProductCrossSellingCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/product-cross-selling", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(ProductCrossSellingCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ProductCrossSellingRepository) SearchAll(ctx ApiContext, criteria Criteria) (*ProductCrossSellingCollection, *http.Response, error) {
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

func (t ProductCrossSellingRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/product-cross-selling", criteria)

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

func (t ProductCrossSellingRepository) Upsert(ctx ApiContext, entity []ProductCrossSelling) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_cross_selling": {
		Entity:  "product_cross_selling",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t ProductCrossSellingRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_cross_selling": {
		Entity:  "product_cross_selling",
		Action:  "delete",
		Payload: payload,
	}})
}

type ProductCrossSelling struct {

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	Id      string  `json:"id,omitempty"`

	SortBy      string  `json:"sortBy,omitempty"`

	Limit      float64  `json:"limit,omitempty"`

	Product      *Product  `json:"product,omitempty"`

	ProductStream      *ProductStream  `json:"productStream,omitempty"`

	ProductStreamId      string  `json:"productStreamId,omitempty"`

	AssignedProducts      []ProductCrossSellingAssignedProducts  `json:"assignedProducts,omitempty"`

	Name      string  `json:"name,omitempty"`

	Position      float64  `json:"position,omitempty"`

	SortDirection      string  `json:"sortDirection,omitempty"`

	ProductId      string  `json:"productId,omitempty"`

	ProductVersionId      string  `json:"productVersionId,omitempty"`

	Type      string  `json:"type,omitempty"`

	Active      bool  `json:"active,omitempty"`

	Translations      []ProductCrossSellingTranslation  `json:"translations,omitempty"`

}

type ProductCrossSellingCollection struct {
	EntityCollection

	Data []ProductCrossSelling `json:"data"`
}

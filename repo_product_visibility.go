package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type ProductVisibilityRepository ClientService

func (t ProductVisibilityRepository) Search(ctx ApiContext, criteria Criteria) (*ProductVisibilityCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/product-visibility", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(ProductVisibilityCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ProductVisibilityRepository) SearchAll(ctx ApiContext, criteria Criteria) (*ProductVisibilityCollection, *http.Response, error) {
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

func (t ProductVisibilityRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/product-visibility", criteria)

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

func (t ProductVisibilityRepository) Upsert(ctx ApiContext, entity []ProductVisibility) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_visibility": {
		Entity:  "product_visibility",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t ProductVisibilityRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_visibility": {
		Entity:  "product_visibility",
		Action:  "delete",
		Payload: payload,
	}})
}

type ProductVisibility struct {

	ProductVersionId      string  `json:"productVersionId,omitempty"`

	SalesChannelId      string  `json:"salesChannelId,omitempty"`

	SalesChannel      *SalesChannel  `json:"salesChannel,omitempty"`

	Id      string  `json:"id,omitempty"`

	ProductId      string  `json:"productId,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Visibility      float64  `json:"visibility,omitempty"`

	Product      *Product  `json:"product,omitempty"`

}

type ProductVisibilityCollection struct {
	EntityCollection

	Data []ProductVisibility `json:"data"`
}

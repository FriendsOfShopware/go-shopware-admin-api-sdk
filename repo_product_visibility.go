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
	SalesChannel *SalesChannel `json:"salesChannel,omitempty"`

	Product *Product `json:"product,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Id string `json:"id,omitempty"`

	ProductVersionId string `json:"productVersionId,omitempty"`

	Visibility float64 `json:"visibility,omitempty"`

	ProductId string `json:"productId,omitempty"`

	SalesChannelId string `json:"salesChannelId,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`
}

type ProductVisibilityCollection struct {
	EntityCollection

	Data []ProductVisibility `json:"data"`
}

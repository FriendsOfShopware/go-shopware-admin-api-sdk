package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type CustomerWishlistRepository ClientService

func (t CustomerWishlistRepository) Search(ctx ApiContext, criteria Criteria) (*CustomerWishlistCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/customer-wishlist", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(CustomerWishlistCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t CustomerWishlistRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/customer-wishlist", criteria)

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

func (t CustomerWishlistRepository) Upsert(ctx ApiContext, entity []CustomerWishlist) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"customer_wishlist": {
		Entity:  "customer_wishlist",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t CustomerWishlistRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"customer_wishlist": {
		Entity:  "customer_wishlist",
		Action:  "delete",
		Payload: payload,
	}})
}

type CustomerWishlist struct {
	CustomFields interface{} `json:"customFields,omitempty"`

	Products []CustomerWishlistProduct `json:"products,omitempty"`

	Customer *Customer `json:"customer,omitempty"`

	SalesChannel *SalesChannel `json:"salesChannel,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Id string `json:"id,omitempty"`

	CustomerId string `json:"customerId,omitempty"`

	SalesChannelId string `json:"salesChannelId,omitempty"`
}

type CustomerWishlistCollection struct {
	EntityCollection

	Data []CustomerWishlist `json:"data"`
}

package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type CustomerWishlistProductRepository ClientService

func (t CustomerWishlistProductRepository) Search(ctx ApiContext, criteria Criteria) (*CustomerWishlistProductCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/customer-wishlist-product", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(CustomerWishlistProductCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t CustomerWishlistProductRepository) SearchAll(ctx ApiContext, criteria Criteria) (*CustomerWishlistProductCollection, *http.Response, error) {
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

func (t CustomerWishlistProductRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/customer-wishlist-product", criteria)

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

func (t CustomerWishlistProductRepository) Upsert(ctx ApiContext, entity []CustomerWishlistProduct) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"customer_wishlist_product": {
		Entity:  "customer_wishlist_product",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t CustomerWishlistProductRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"customer_wishlist_product": {
		Entity:  "customer_wishlist_product",
		Action:  "delete",
		Payload: payload,
	}})
}

type CustomerWishlistProduct struct {

	WishlistId      string  `json:"wishlistId,omitempty"`

	Wishlist      *CustomerWishlist  `json:"wishlist,omitempty"`

	Product      *Product  `json:"product,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Id      string  `json:"id,omitempty"`

	ProductId      string  `json:"productId,omitempty"`

	ProductVersionId      string  `json:"productVersionId,omitempty"`

}

type CustomerWishlistProductCollection struct {
	EntityCollection

	Data []CustomerWishlistProduct `json:"data"`
}

package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type CustomerWishlistProductRepository struct {
	*GenericRepository[CustomerWishlistProduct]
}

func NewCustomerWishlistProductRepository(client *Client) *CustomerWishlistProductRepository {
	return &CustomerWishlistProductRepository{
		GenericRepository: NewGenericRepository[CustomerWishlistProduct](client),
	}
}

func (t *CustomerWishlistProductRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[CustomerWishlistProduct], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "customer-wishlist-product")
}

func (t *CustomerWishlistProductRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[CustomerWishlistProduct], *http.Response, error) {
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

func (t *CustomerWishlistProductRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "customer-wishlist-product")
}

func (t *CustomerWishlistProductRepository) Upsert(ctx ApiContext, entity []CustomerWishlistProduct) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "customer_wishlist_product")
}

func (t *CustomerWishlistProductRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "customer_wishlist_product")
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

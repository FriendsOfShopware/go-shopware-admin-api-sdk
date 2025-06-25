package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type CustomerWishlistRepository struct {
	*GenericRepository[CustomerWishlist]
}

func NewCustomerWishlistRepository(client *Client) *CustomerWishlistRepository {
	return &CustomerWishlistRepository{
		GenericRepository: NewGenericRepository[CustomerWishlist](client),
	}
}

func (t *CustomerWishlistRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[CustomerWishlist], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "customer-wishlist")
}

func (t *CustomerWishlistRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[CustomerWishlist], *http.Response, error) {
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

func (t *CustomerWishlistRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "customer-wishlist")
}

func (t *CustomerWishlistRepository) Upsert(ctx ApiContext, entity []CustomerWishlist) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "customer_wishlist")
}

func (t *CustomerWishlistRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "customer_wishlist")
}

type CustomerWishlist struct {

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Products      []CustomerWishlistProduct  `json:"products,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Id      string  `json:"id,omitempty"`

	CustomerId      string  `json:"customerId,omitempty"`

	SalesChannelId      string  `json:"salesChannelId,omitempty"`

	Customer      *Customer  `json:"customer,omitempty"`

	SalesChannel      *SalesChannel  `json:"salesChannel,omitempty"`

}

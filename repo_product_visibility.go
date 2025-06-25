package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type ProductVisibilityRepository struct {
	*GenericRepository[ProductVisibility]
}

func NewProductVisibilityRepository(client *Client) *ProductVisibilityRepository {
	return &ProductVisibilityRepository{
		GenericRepository: NewGenericRepository[ProductVisibility](client),
	}
}

func (t *ProductVisibilityRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductVisibility], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "product-visibility")
}

func (t *ProductVisibilityRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductVisibility], *http.Response, error) {
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

func (t *ProductVisibilityRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "product-visibility")
}

func (t *ProductVisibilityRepository) Upsert(ctx ApiContext, entity []ProductVisibility) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "product_visibility")
}

func (t *ProductVisibilityRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "product_visibility")
}

type ProductVisibility struct {

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Id      string  `json:"id,omitempty"`

	Product      *Product  `json:"product,omitempty"`

	ProductId      string  `json:"productId,omitempty"`

	ProductVersionId      string  `json:"productVersionId,omitempty"`

	SalesChannel      *SalesChannel  `json:"salesChannel,omitempty"`

	SalesChannelId      string  `json:"salesChannelId,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Visibility      float64  `json:"visibility,omitempty"`

}

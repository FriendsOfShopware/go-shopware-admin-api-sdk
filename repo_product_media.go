package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type ProductMediaRepository struct {
	*GenericRepository[ProductMedia]
}

func NewProductMediaRepository(client *Client) *ProductMediaRepository {
	return &ProductMediaRepository{
		GenericRepository: NewGenericRepository[ProductMedia](client),
	}
}

func (t *ProductMediaRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductMedia], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "product-media")
}

func (t *ProductMediaRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductMedia], *http.Response, error) {
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

func (t *ProductMediaRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "product-media")
}

func (t *ProductMediaRepository) Upsert(ctx ApiContext, entity []ProductMedia) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "product_media")
}

func (t *ProductMediaRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "product_media")
}

type ProductMedia struct {

	CoverProducts      []Product  `json:"coverProducts,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Id      string  `json:"id,omitempty"`

	Media      *Media  `json:"media,omitempty"`

	MediaId      string  `json:"mediaId,omitempty"`

	Position      float64  `json:"position,omitempty"`

	Product      *Product  `json:"product,omitempty"`

	ProductId      string  `json:"productId,omitempty"`

	ProductVersionId      string  `json:"productVersionId,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	VersionId      string  `json:"versionId,omitempty"`

}

package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type ProductCrossSellingRepository struct {
	*GenericRepository[ProductCrossSelling]
}

func NewProductCrossSellingRepository(client *Client) *ProductCrossSellingRepository {
	return &ProductCrossSellingRepository{
		GenericRepository: NewGenericRepository[ProductCrossSelling](client),
	}
}

func (t *ProductCrossSellingRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductCrossSelling], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "product-cross-selling")
}

func (t *ProductCrossSellingRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductCrossSelling], *http.Response, error) {
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

func (t *ProductCrossSellingRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "product-cross-selling")
}

func (t *ProductCrossSellingRepository) Upsert(ctx ApiContext, entity []ProductCrossSelling) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "product_cross_selling")
}

func (t *ProductCrossSellingRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "product_cross_selling")
}

type ProductCrossSelling struct {

	SortBy      string  `json:"sortBy,omitempty"`

	SortDirection      string  `json:"sortDirection,omitempty"`

	Type      string  `json:"type,omitempty"`

	Active      bool  `json:"active,omitempty"`

	Limit      float64  `json:"limit,omitempty"`

	ProductId      string  `json:"productId,omitempty"`

	ProductStreamId      string  `json:"productStreamId,omitempty"`

	ProductStream      *ProductStream  `json:"productStream,omitempty"`

	Product      *Product  `json:"product,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Name      string  `json:"name,omitempty"`

	Position      float64  `json:"position,omitempty"`

	Translations      []ProductCrossSellingTranslation  `json:"translations,omitempty"`

	Id      string  `json:"id,omitempty"`

	ProductVersionId      string  `json:"productVersionId,omitempty"`

	AssignedProducts      []ProductCrossSellingAssignedProducts  `json:"assignedProducts,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

}

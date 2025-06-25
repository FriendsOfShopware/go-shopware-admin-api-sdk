package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type ProductSortingRepository struct {
	*GenericRepository[ProductSorting]
}

func NewProductSortingRepository(client *Client) *ProductSortingRepository {
	return &ProductSortingRepository{
		GenericRepository: NewGenericRepository[ProductSorting](client),
	}
}

func (t *ProductSortingRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductSorting], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "product-sorting")
}

func (t *ProductSortingRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductSorting], *http.Response, error) {
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

func (t *ProductSortingRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "product-sorting")
}

func (t *ProductSortingRepository) Upsert(ctx ApiContext, entity []ProductSorting) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "product_sorting")
}

func (t *ProductSortingRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "product_sorting")
}

type ProductSorting struct {

	Active      bool  `json:"active,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Fields      interface{}  `json:"fields,omitempty"`

	Id      string  `json:"id,omitempty"`

	Key      string  `json:"key,omitempty"`

	Label      string  `json:"label,omitempty"`

	Locked      bool  `json:"locked,omitempty"`

	Priority      float64  `json:"priority,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	Translations      []ProductSortingTranslation  `json:"translations,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}

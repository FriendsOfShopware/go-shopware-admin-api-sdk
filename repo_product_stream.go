package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type ProductStreamRepository struct {
	*GenericRepository[ProductStream]
}

func NewProductStreamRepository(client *Client) *ProductStreamRepository {
	return &ProductStreamRepository{
		GenericRepository: NewGenericRepository[ProductStream](client),
	}
}

func (t *ProductStreamRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductStream], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "product-stream")
}

func (t *ProductStreamRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductStream], *http.Response, error) {
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

func (t *ProductStreamRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "product-stream")
}

func (t *ProductStreamRepository) Upsert(ctx ApiContext, entity []ProductStream) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "product_stream")
}

func (t *ProductStreamRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "product_stream")
}

type ProductStream struct {

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	Id      string  `json:"id,omitempty"`

	ApiFilter      interface{}  `json:"apiFilter,omitempty"`

	Invalid      bool  `json:"invalid,omitempty"`

	Name      string  `json:"name,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Translations      []ProductStreamTranslation  `json:"translations,omitempty"`

	Filters      []ProductStreamFilter  `json:"filters,omitempty"`

	ProductExports      []ProductExport  `json:"productExports,omitempty"`

	Description      string  `json:"description,omitempty"`

	ProductCrossSellings      []ProductCrossSelling  `json:"productCrossSellings,omitempty"`

	Categories      []Category  `json:"categories,omitempty"`

}

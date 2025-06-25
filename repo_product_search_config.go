package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type ProductSearchConfigRepository struct {
	*GenericRepository[ProductSearchConfig]
}

func NewProductSearchConfigRepository(client *Client) *ProductSearchConfigRepository {
	return &ProductSearchConfigRepository{
		GenericRepository: NewGenericRepository[ProductSearchConfig](client),
	}
}

func (t *ProductSearchConfigRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductSearchConfig], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "product-search-config")
}

func (t *ProductSearchConfigRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductSearchConfig], *http.Response, error) {
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

func (t *ProductSearchConfigRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "product-search-config")
}

func (t *ProductSearchConfigRepository) Upsert(ctx ApiContext, entity []ProductSearchConfig) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "product_search_config")
}

func (t *ProductSearchConfigRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "product_search_config")
}

type ProductSearchConfig struct {

	ConfigFields      []ProductSearchConfigField  `json:"configFields,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Id      string  `json:"id,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	AndLogic      bool  `json:"andLogic,omitempty"`

	MinSearchLength      float64  `json:"minSearchLength,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	ExcludedTerms      interface{}  `json:"excludedTerms,omitempty"`

	Language      *Language  `json:"language,omitempty"`

}

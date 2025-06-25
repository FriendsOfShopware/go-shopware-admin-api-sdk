package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type ProductSearchConfigFieldRepository struct {
	*GenericRepository[ProductSearchConfigField]
}

func NewProductSearchConfigFieldRepository(client *Client) *ProductSearchConfigFieldRepository {
	return &ProductSearchConfigFieldRepository{
		GenericRepository: NewGenericRepository[ProductSearchConfigField](client),
	}
}

func (t *ProductSearchConfigFieldRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductSearchConfigField], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "product-search-config-field")
}

func (t *ProductSearchConfigFieldRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductSearchConfigField], *http.Response, error) {
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

func (t *ProductSearchConfigFieldRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "product-search-config-field")
}

func (t *ProductSearchConfigFieldRepository) Upsert(ctx ApiContext, entity []ProductSearchConfigField) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "product_search_config_field")
}

func (t *ProductSearchConfigFieldRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "product_search_config_field")
}

type ProductSearchConfigField struct {

	Id      string  `json:"id,omitempty"`

	Tokenize      bool  `json:"tokenize,omitempty"`

	Ranking      float64  `json:"ranking,omitempty"`

	SearchConfig      *ProductSearchConfig  `json:"searchConfig,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	SearchConfigId      string  `json:"searchConfigId,omitempty"`

	CustomFieldId      string  `json:"customFieldId,omitempty"`

	Field      string  `json:"field,omitempty"`

	Searchable      bool  `json:"searchable,omitempty"`

	CustomField      *CustomField  `json:"customField,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

}

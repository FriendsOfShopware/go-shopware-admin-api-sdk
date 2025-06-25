package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type ProductFeatureSetRepository struct {
	*GenericRepository[ProductFeatureSet]
}

func NewProductFeatureSetRepository(client *Client) *ProductFeatureSetRepository {
	return &ProductFeatureSetRepository{
		GenericRepository: NewGenericRepository[ProductFeatureSet](client),
	}
}

func (t *ProductFeatureSetRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductFeatureSet], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "product-feature-set")
}

func (t *ProductFeatureSetRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductFeatureSet], *http.Response, error) {
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

func (t *ProductFeatureSetRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "product-feature-set")
}

func (t *ProductFeatureSetRepository) Upsert(ctx ApiContext, entity []ProductFeatureSet) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "product_feature_set")
}

func (t *ProductFeatureSetRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "product_feature_set")
}

type ProductFeatureSet struct {

	Translations      []ProductFeatureSetTranslation  `json:"translations,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Id      string  `json:"id,omitempty"`

	Description      string  `json:"description,omitempty"`

	Products      []Product  `json:"products,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	Name      string  `json:"name,omitempty"`

	Features      interface{}  `json:"features,omitempty"`

}

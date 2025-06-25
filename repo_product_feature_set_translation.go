package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type ProductFeatureSetTranslationRepository struct {
	*GenericRepository[ProductFeatureSetTranslation]
}

func NewProductFeatureSetTranslationRepository(client *Client) *ProductFeatureSetTranslationRepository {
	return &ProductFeatureSetTranslationRepository{
		GenericRepository: NewGenericRepository[ProductFeatureSetTranslation](client),
	}
}

func (t *ProductFeatureSetTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductFeatureSetTranslation], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "product-feature-set-translation")
}

func (t *ProductFeatureSetTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductFeatureSetTranslation], *http.Response, error) {
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

func (t *ProductFeatureSetTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "product-feature-set-translation")
}

func (t *ProductFeatureSetTranslationRepository) Upsert(ctx ApiContext, entity []ProductFeatureSetTranslation) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "product_feature_set_translation")
}

func (t *ProductFeatureSetTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "product_feature_set_translation")
}

type ProductFeatureSetTranslation struct {

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Description      string  `json:"description,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	Name      string  `json:"name,omitempty"`

	ProductFeatureSet      *ProductFeatureSet  `json:"productFeatureSet,omitempty"`

	ProductFeatureSetId      string  `json:"productFeatureSetId,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}

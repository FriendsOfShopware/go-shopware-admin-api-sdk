package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type ProductCrossSellingTranslationRepository struct {
	*GenericRepository[ProductCrossSellingTranslation]
}

func NewProductCrossSellingTranslationRepository(client *Client) *ProductCrossSellingTranslationRepository {
	return &ProductCrossSellingTranslationRepository{
		GenericRepository: NewGenericRepository[ProductCrossSellingTranslation](client),
	}
}

func (t *ProductCrossSellingTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductCrossSellingTranslation], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "product-cross-selling-translation")
}

func (t *ProductCrossSellingTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductCrossSellingTranslation], *http.Response, error) {
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

func (t *ProductCrossSellingTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "product-cross-selling-translation")
}

func (t *ProductCrossSellingTranslationRepository) Upsert(ctx ApiContext, entity []ProductCrossSellingTranslation) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "product_cross_selling_translation")
}

func (t *ProductCrossSellingTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "product_cross_selling_translation")
}

type ProductCrossSellingTranslation struct {

	ProductCrossSellingId      string  `json:"productCrossSellingId,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	ProductCrossSelling      *ProductCrossSelling  `json:"productCrossSelling,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	Name      string  `json:"name,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}

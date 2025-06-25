package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type ProductSortingTranslationRepository struct {
	*GenericRepository[ProductSortingTranslation]
}

func NewProductSortingTranslationRepository(client *Client) *ProductSortingTranslationRepository {
	return &ProductSortingTranslationRepository{
		GenericRepository: NewGenericRepository[ProductSortingTranslation](client),
	}
}

func (t *ProductSortingTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductSortingTranslation], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "product-sorting-translation")
}

func (t *ProductSortingTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductSortingTranslation], *http.Response, error) {
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

func (t *ProductSortingTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "product-sorting-translation")
}

func (t *ProductSortingTranslationRepository) Upsert(ctx ApiContext, entity []ProductSortingTranslation) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "product_sorting_translation")
}

func (t *ProductSortingTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "product_sorting_translation")
}

type ProductSortingTranslation struct {

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	ProductSortingId      string  `json:"productSortingId,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	ProductSorting      *ProductSorting  `json:"productSorting,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	Label      string  `json:"label,omitempty"`

}

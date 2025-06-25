package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type ProductTranslationRepository struct {
	*GenericRepository[ProductTranslation]
}

func NewProductTranslationRepository(client *Client) *ProductTranslationRepository {
	return &ProductTranslationRepository{
		GenericRepository: NewGenericRepository[ProductTranslation](client),
	}
}

func (t *ProductTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductTranslation], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "product-translation")
}

func (t *ProductTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductTranslation], *http.Response, error) {
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

func (t *ProductTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "product-translation")
}

func (t *ProductTranslationRepository) Upsert(ctx ApiContext, entity []ProductTranslation) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "product_translation")
}

func (t *ProductTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "product_translation")
}

type ProductTranslation struct {

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	CustomSearchKeywords      interface{}  `json:"customSearchKeywords,omitempty"`

	Description      string  `json:"description,omitempty"`

	Keywords      string  `json:"keywords,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	MetaDescription      string  `json:"metaDescription,omitempty"`

	MetaTitle      string  `json:"metaTitle,omitempty"`

	Name      string  `json:"name,omitempty"`

	PackUnit      string  `json:"packUnit,omitempty"`

	PackUnitPlural      string  `json:"packUnitPlural,omitempty"`

	Product      *Product  `json:"product,omitempty"`

	ProductId      string  `json:"productId,omitempty"`

	ProductVersionId      string  `json:"productVersionId,omitempty"`

	SlotConfig      interface{}  `json:"slotConfig,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}

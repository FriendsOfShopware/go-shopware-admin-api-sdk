package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type ProductManufacturerTranslationRepository struct {
	*GenericRepository[ProductManufacturerTranslation]
}

func NewProductManufacturerTranslationRepository(client *Client) *ProductManufacturerTranslationRepository {
	return &ProductManufacturerTranslationRepository{
		GenericRepository: NewGenericRepository[ProductManufacturerTranslation](client),
	}
}

func (t *ProductManufacturerTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductManufacturerTranslation], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "product-manufacturer-translation")
}

func (t *ProductManufacturerTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductManufacturerTranslation], *http.Response, error) {
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

func (t *ProductManufacturerTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "product-manufacturer-translation")
}

func (t *ProductManufacturerTranslationRepository) Upsert(ctx ApiContext, entity []ProductManufacturerTranslation) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "product_manufacturer_translation")
}

func (t *ProductManufacturerTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "product_manufacturer_translation")
}

type ProductManufacturerTranslation struct {

	ProductManufacturerId      string  `json:"productManufacturerId,omitempty"`

	ProductManufacturerVersionId      string  `json:"productManufacturerVersionId,omitempty"`

	Name      string  `json:"name,omitempty"`

	Description      string  `json:"description,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	ProductManufacturer      *ProductManufacturer  `json:"productManufacturer,omitempty"`

	Language      *Language  `json:"language,omitempty"`

}

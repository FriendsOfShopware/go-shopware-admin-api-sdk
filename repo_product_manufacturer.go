package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type ProductManufacturerRepository struct {
	*GenericRepository[ProductManufacturer]
}

func NewProductManufacturerRepository(client *Client) *ProductManufacturerRepository {
	return &ProductManufacturerRepository{
		GenericRepository: NewGenericRepository[ProductManufacturer](client),
	}
}

func (t *ProductManufacturerRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductManufacturer], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "product-manufacturer")
}

func (t *ProductManufacturerRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductManufacturer], *http.Response, error) {
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

func (t *ProductManufacturerRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "product-manufacturer")
}

func (t *ProductManufacturerRepository) Upsert(ctx ApiContext, entity []ProductManufacturer) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "product_manufacturer")
}

func (t *ProductManufacturerRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "product_manufacturer")
}

type ProductManufacturer struct {

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	VersionId      string  `json:"versionId,omitempty"`

	Description      string  `json:"description,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Translations      []ProductManufacturerTranslation  `json:"translations,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	Id      string  `json:"id,omitempty"`

	MediaId      string  `json:"mediaId,omitempty"`

	Link      string  `json:"link,omitempty"`

	Name      string  `json:"name,omitempty"`

	Media      *Media  `json:"media,omitempty"`

	Products      []Product  `json:"products,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

}

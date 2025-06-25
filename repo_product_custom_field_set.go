package go_shopware_admin_sdk

import (
	"net/http"

)

type ProductCustomFieldSetRepository struct {
	*GenericRepository[ProductCustomFieldSet]
}

func NewProductCustomFieldSetRepository(client *Client) *ProductCustomFieldSetRepository {
	return &ProductCustomFieldSetRepository{
		GenericRepository: NewGenericRepository[ProductCustomFieldSet](client),
	}
}

func (t *ProductCustomFieldSetRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductCustomFieldSet], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "product-custom-field-set")
}

func (t *ProductCustomFieldSetRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductCustomFieldSet], *http.Response, error) {
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

func (t *ProductCustomFieldSetRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "product-custom-field-set")
}

func (t *ProductCustomFieldSetRepository) Upsert(ctx ApiContext, entity []ProductCustomFieldSet) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "product_custom_field_set")
}

func (t *ProductCustomFieldSetRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "product_custom_field_set")
}

type ProductCustomFieldSet struct {

	ProductId      string  `json:"productId,omitempty"`

	CustomFieldSetId      string  `json:"customFieldSetId,omitempty"`

	ProductVersionId      string  `json:"productVersionId,omitempty"`

	Product      *Product  `json:"product,omitempty"`

	CustomFieldSet      *CustomFieldSet  `json:"customFieldSet,omitempty"`

}

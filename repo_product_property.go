package go_shopware_admin_sdk

import (
	"net/http"

)

type ProductPropertyRepository struct {
	*GenericRepository[ProductProperty]
}

func NewProductPropertyRepository(client *Client) *ProductPropertyRepository {
	return &ProductPropertyRepository{
		GenericRepository: NewGenericRepository[ProductProperty](client),
	}
}

func (t *ProductPropertyRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductProperty], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "product-property")
}

func (t *ProductPropertyRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductProperty], *http.Response, error) {
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

func (t *ProductPropertyRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "product-property")
}

func (t *ProductPropertyRepository) Upsert(ctx ApiContext, entity []ProductProperty) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "product_property")
}

func (t *ProductPropertyRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "product_property")
}

type ProductProperty struct {

	Option      *PropertyGroupOption  `json:"option,omitempty"`

	OptionId      string  `json:"optionId,omitempty"`

	Product      *Product  `json:"product,omitempty"`

	ProductId      string  `json:"productId,omitempty"`

	ProductVersionId      string  `json:"productVersionId,omitempty"`

}

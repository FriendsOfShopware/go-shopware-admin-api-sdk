package go_shopware_admin_sdk

import (
	"net/http"

)

type ProductOptionRepository struct {
	*GenericRepository[ProductOption]
}

func NewProductOptionRepository(client *Client) *ProductOptionRepository {
	return &ProductOptionRepository{
		GenericRepository: NewGenericRepository[ProductOption](client),
	}
}

func (t *ProductOptionRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductOption], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "product-option")
}

func (t *ProductOptionRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductOption], *http.Response, error) {
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

func (t *ProductOptionRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "product-option")
}

func (t *ProductOptionRepository) Upsert(ctx ApiContext, entity []ProductOption) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "product_option")
}

func (t *ProductOptionRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "product_option")
}

type ProductOption struct {

	Option      *PropertyGroupOption  `json:"option,omitempty"`

	OptionId      string  `json:"optionId,omitempty"`

	Product      *Product  `json:"product,omitempty"`

	ProductId      string  `json:"productId,omitempty"`

	ProductVersionId      string  `json:"productVersionId,omitempty"`

}

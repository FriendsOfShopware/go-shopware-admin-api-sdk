package go_shopware_admin_sdk

import (
	"net/http"

)

type ProductCategoryRepository struct {
	*GenericRepository[ProductCategory]
}

func NewProductCategoryRepository(client *Client) *ProductCategoryRepository {
	return &ProductCategoryRepository{
		GenericRepository: NewGenericRepository[ProductCategory](client),
	}
}

func (t *ProductCategoryRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductCategory], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "product-category")
}

func (t *ProductCategoryRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductCategory], *http.Response, error) {
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

func (t *ProductCategoryRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "product-category")
}

func (t *ProductCategoryRepository) Upsert(ctx ApiContext, entity []ProductCategory) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "product_category")
}

func (t *ProductCategoryRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "product_category")
}

type ProductCategory struct {

	ProductVersionId      string  `json:"productVersionId,omitempty"`

	CategoryId      string  `json:"categoryId,omitempty"`

	CategoryVersionId      string  `json:"categoryVersionId,omitempty"`

	Product      *Product  `json:"product,omitempty"`

	Category      *Category  `json:"category,omitempty"`

	ProductId      string  `json:"productId,omitempty"`

}

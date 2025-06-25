package go_shopware_admin_sdk

import (
	"net/http"

)

type ProductCategoryTreeRepository struct {
	*GenericRepository[ProductCategoryTree]
}

func NewProductCategoryTreeRepository(client *Client) *ProductCategoryTreeRepository {
	return &ProductCategoryTreeRepository{
		GenericRepository: NewGenericRepository[ProductCategoryTree](client),
	}
}

func (t *ProductCategoryTreeRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductCategoryTree], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "product-category-tree")
}

func (t *ProductCategoryTreeRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductCategoryTree], *http.Response, error) {
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

func (t *ProductCategoryTreeRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "product-category-tree")
}

func (t *ProductCategoryTreeRepository) Upsert(ctx ApiContext, entity []ProductCategoryTree) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "product_category_tree")
}

func (t *ProductCategoryTreeRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "product_category_tree")
}

type ProductCategoryTree struct {

	ProductId      string  `json:"productId,omitempty"`

	ProductVersionId      string  `json:"productVersionId,omitempty"`

	CategoryId      string  `json:"categoryId,omitempty"`

	CategoryVersionId      string  `json:"categoryVersionId,omitempty"`

	Product      *Product  `json:"product,omitempty"`

	Category      *Category  `json:"category,omitempty"`

}

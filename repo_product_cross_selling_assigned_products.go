package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type ProductCrossSellingAssignedProductsRepository struct {
	*GenericRepository[ProductCrossSellingAssignedProducts]
}

func NewProductCrossSellingAssignedProductsRepository(client *Client) *ProductCrossSellingAssignedProductsRepository {
	return &ProductCrossSellingAssignedProductsRepository{
		GenericRepository: NewGenericRepository[ProductCrossSellingAssignedProducts](client),
	}
}

func (t *ProductCrossSellingAssignedProductsRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductCrossSellingAssignedProducts], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "product-cross-selling-assigned-products")
}

func (t *ProductCrossSellingAssignedProductsRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductCrossSellingAssignedProducts], *http.Response, error) {
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

func (t *ProductCrossSellingAssignedProductsRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "product-cross-selling-assigned-products")
}

func (t *ProductCrossSellingAssignedProductsRepository) Upsert(ctx ApiContext, entity []ProductCrossSellingAssignedProducts) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "product_cross_selling_assigned_products")
}

func (t *ProductCrossSellingAssignedProductsRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "product_cross_selling_assigned_products")
}

type ProductCrossSellingAssignedProducts struct {

	Id      string  `json:"id,omitempty"`

	ProductVersionId      string  `json:"productVersionId,omitempty"`

	Product      *Product  `json:"product,omitempty"`

	CrossSelling      *ProductCrossSelling  `json:"crossSelling,omitempty"`

	Position      float64  `json:"position,omitempty"`

	CrossSellingId      string  `json:"crossSellingId,omitempty"`

	ProductId      string  `json:"productId,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}

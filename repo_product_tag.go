package go_shopware_admin_sdk

import (
	"net/http"

)

type ProductTagRepository struct {
	*GenericRepository[ProductTag]
}

func NewProductTagRepository(client *Client) *ProductTagRepository {
	return &ProductTagRepository{
		GenericRepository: NewGenericRepository[ProductTag](client),
	}
}

func (t *ProductTagRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductTag], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "product-tag")
}

func (t *ProductTagRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductTag], *http.Response, error) {
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

func (t *ProductTagRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "product-tag")
}

func (t *ProductTagRepository) Upsert(ctx ApiContext, entity []ProductTag) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "product_tag")
}

func (t *ProductTagRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "product_tag")
}

type ProductTag struct {

	Product      *Product  `json:"product,omitempty"`

	ProductId      string  `json:"productId,omitempty"`

	ProductVersionId      string  `json:"productVersionId,omitempty"`

	Tag      *Tag  `json:"tag,omitempty"`

	TagId      string  `json:"tagId,omitempty"`

}

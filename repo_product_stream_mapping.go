package go_shopware_admin_sdk

import (
	"net/http"

)

type ProductStreamMappingRepository struct {
	*GenericRepository[ProductStreamMapping]
}

func NewProductStreamMappingRepository(client *Client) *ProductStreamMappingRepository {
	return &ProductStreamMappingRepository{
		GenericRepository: NewGenericRepository[ProductStreamMapping](client),
	}
}

func (t *ProductStreamMappingRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductStreamMapping], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "product-stream-mapping")
}

func (t *ProductStreamMappingRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductStreamMapping], *http.Response, error) {
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

func (t *ProductStreamMappingRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "product-stream-mapping")
}

func (t *ProductStreamMappingRepository) Upsert(ctx ApiContext, entity []ProductStreamMapping) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "product_stream_mapping")
}

func (t *ProductStreamMappingRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "product_stream_mapping")
}

type ProductStreamMapping struct {

	ProductId      string  `json:"productId,omitempty"`

	ProductVersionId      string  `json:"productVersionId,omitempty"`

	ProductStreamId      string  `json:"productStreamId,omitempty"`

	Product      *Product  `json:"product,omitempty"`

	ProductStream      *ProductStream  `json:"productStream,omitempty"`

}

package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type ProductStreamFilterRepository struct {
	*GenericRepository[ProductStreamFilter]
}

func NewProductStreamFilterRepository(client *Client) *ProductStreamFilterRepository {
	return &ProductStreamFilterRepository{
		GenericRepository: NewGenericRepository[ProductStreamFilter](client),
	}
}

func (t *ProductStreamFilterRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductStreamFilter], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "product-stream-filter")
}

func (t *ProductStreamFilterRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductStreamFilter], *http.Response, error) {
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

func (t *ProductStreamFilterRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "product-stream-filter")
}

func (t *ProductStreamFilterRepository) Upsert(ctx ApiContext, entity []ProductStreamFilter) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "product_stream_filter")
}

func (t *ProductStreamFilterRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "product_stream_filter")
}

type ProductStreamFilter struct {

	Type      string  `json:"type,omitempty"`

	Value      string  `json:"value,omitempty"`

	ProductStream      *ProductStream  `json:"productStream,omitempty"`

	Parent      *ProductStreamFilter  `json:"parent,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Id      string  `json:"id,omitempty"`

	ParentId      string  `json:"parentId,omitempty"`

	Operator      string  `json:"operator,omitempty"`

	Parameters      interface{}  `json:"parameters,omitempty"`

	ProductStreamId      string  `json:"productStreamId,omitempty"`

	Field      string  `json:"field,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Position      float64  `json:"position,omitempty"`

	Queries      []ProductStreamFilter  `json:"queries,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

}

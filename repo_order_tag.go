package go_shopware_admin_sdk

import (
	"net/http"

)

type OrderTagRepository struct {
	*GenericRepository[OrderTag]
}

func NewOrderTagRepository(client *Client) *OrderTagRepository {
	return &OrderTagRepository{
		GenericRepository: NewGenericRepository[OrderTag](client),
	}
}

func (t *OrderTagRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[OrderTag], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "order-tag")
}

func (t *OrderTagRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[OrderTag], *http.Response, error) {
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

func (t *OrderTagRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "order-tag")
}

func (t *OrderTagRepository) Upsert(ctx ApiContext, entity []OrderTag) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "order_tag")
}

func (t *OrderTagRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "order_tag")
}

type OrderTag struct {

	OrderId      string  `json:"orderId,omitempty"`

	OrderVersionId      string  `json:"orderVersionId,omitempty"`

	TagId      string  `json:"tagId,omitempty"`

	Order      *Order  `json:"order,omitempty"`

	Tag      *Tag  `json:"tag,omitempty"`

}

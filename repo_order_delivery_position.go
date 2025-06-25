package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type OrderDeliveryPositionRepository struct {
	*GenericRepository[OrderDeliveryPosition]
}

func NewOrderDeliveryPositionRepository(client *Client) *OrderDeliveryPositionRepository {
	return &OrderDeliveryPositionRepository{
		GenericRepository: NewGenericRepository[OrderDeliveryPosition](client),
	}
}

func (t *OrderDeliveryPositionRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[OrderDeliveryPosition], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "order-delivery-position")
}

func (t *OrderDeliveryPositionRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[OrderDeliveryPosition], *http.Response, error) {
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

func (t *OrderDeliveryPositionRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "order-delivery-position")
}

func (t *OrderDeliveryPositionRepository) Upsert(ctx ApiContext, entity []OrderDeliveryPosition) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "order_delivery_position")
}

func (t *OrderDeliveryPositionRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "order_delivery_position")
}

type OrderDeliveryPosition struct {

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Id      string  `json:"id,omitempty"`

	OrderDelivery      *OrderDelivery  `json:"orderDelivery,omitempty"`

	OrderDeliveryId      string  `json:"orderDeliveryId,omitempty"`

	OrderDeliveryVersionId      string  `json:"orderDeliveryVersionId,omitempty"`

	OrderLineItem      *OrderLineItem  `json:"orderLineItem,omitempty"`

	OrderLineItemId      string  `json:"orderLineItemId,omitempty"`

	OrderLineItemVersionId      string  `json:"orderLineItemVersionId,omitempty"`

	Price      interface{}  `json:"price,omitempty"`

	Quantity      float64  `json:"quantity,omitempty"`

	TotalPrice      float64  `json:"totalPrice,omitempty"`

	UnitPrice      float64  `json:"unitPrice,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	VersionId      string  `json:"versionId,omitempty"`

}

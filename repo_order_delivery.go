package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type OrderDeliveryRepository struct {
	*GenericRepository[OrderDelivery]
}

func NewOrderDeliveryRepository(client *Client) *OrderDeliveryRepository {
	return &OrderDeliveryRepository{
		GenericRepository: NewGenericRepository[OrderDelivery](client),
	}
}

func (t *OrderDeliveryRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[OrderDelivery], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "order-delivery")
}

func (t *OrderDeliveryRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[OrderDelivery], *http.Response, error) {
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

func (t *OrderDeliveryRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "order-delivery")
}

func (t *OrderDeliveryRepository) Upsert(ctx ApiContext, entity []OrderDelivery) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "order_delivery")
}

func (t *OrderDeliveryRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "order_delivery")
}

type OrderDelivery struct {

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Id      string  `json:"id,omitempty"`

	Order      *Order  `json:"order,omitempty"`

	OrderId      string  `json:"orderId,omitempty"`

	OrderVersionId      string  `json:"orderVersionId,omitempty"`

	Positions      []OrderDeliveryPosition  `json:"positions,omitempty"`

	ShippingCosts      interface{}  `json:"shippingCosts,omitempty"`

	ShippingDateEarliest      time.Time  `json:"shippingDateEarliest,omitempty"`

	ShippingDateLatest      time.Time  `json:"shippingDateLatest,omitempty"`

	ShippingMethod      *ShippingMethod  `json:"shippingMethod,omitempty"`

	ShippingMethodId      string  `json:"shippingMethodId,omitempty"`

	ShippingOrderAddress      *OrderAddress  `json:"shippingOrderAddress,omitempty"`

	ShippingOrderAddressId      string  `json:"shippingOrderAddressId,omitempty"`

	ShippingOrderAddressVersionId      string  `json:"shippingOrderAddressVersionId,omitempty"`

	StateId      string  `json:"stateId,omitempty"`

	StateMachineState      *StateMachineState  `json:"stateMachineState,omitempty"`

	TrackingCodes      interface{}  `json:"trackingCodes,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	VersionId      string  `json:"versionId,omitempty"`

}

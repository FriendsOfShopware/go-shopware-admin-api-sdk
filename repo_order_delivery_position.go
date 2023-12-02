package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type OrderDeliveryPositionRepository ClientService

func (t OrderDeliveryPositionRepository) Search(ctx ApiContext, criteria Criteria) (*OrderDeliveryPositionCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/order-delivery-position", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(OrderDeliveryPositionCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t OrderDeliveryPositionRepository) SearchAll(ctx ApiContext, criteria Criteria) (*OrderDeliveryPositionCollection, *http.Response, error) {
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

func (t OrderDeliveryPositionRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/order-delivery-position", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(SearchIdsResponse)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t OrderDeliveryPositionRepository) Upsert(ctx ApiContext, entity []OrderDeliveryPosition) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"order_delivery_position": {
		Entity:  "order_delivery_position",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t OrderDeliveryPositionRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"order_delivery_position": {
		Entity:  "order_delivery_position",
		Action:  "delete",
		Payload: payload,
	}})
}

type OrderDeliveryPosition struct {

	Id      string  `json:"id,omitempty"`

	VersionId      string  `json:"versionId,omitempty"`

	OrderLineItemId      string  `json:"orderLineItemId,omitempty"`

	OrderLineItemVersionId      string  `json:"orderLineItemVersionId,omitempty"`

	OrderDelivery      *OrderDelivery  `json:"orderDelivery,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Quantity      float64  `json:"quantity,omitempty"`

	Price      interface{}  `json:"price,omitempty"`

	UnitPrice      float64  `json:"unitPrice,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	OrderLineItem      *OrderLineItem  `json:"orderLineItem,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	OrderDeliveryId      string  `json:"orderDeliveryId,omitempty"`

	OrderDeliveryVersionId      string  `json:"orderDeliveryVersionId,omitempty"`

	TotalPrice      float64  `json:"totalPrice,omitempty"`

}

type OrderDeliveryPositionCollection struct {
	EntityCollection

	Data []OrderDeliveryPosition `json:"data"`
}

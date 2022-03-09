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
	Id string `json:"id,omitempty"`

	OrderDeliveryId string `json:"orderDeliveryId,omitempty"`

	OrderDeliveryVersionId string `json:"orderDeliveryVersionId,omitempty"`

	UnitPrice float64 `json:"unitPrice,omitempty"`

	TotalPrice float64 `json:"totalPrice,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	OrderLineItemVersionId string `json:"orderLineItemVersionId,omitempty"`

	Price interface{} `json:"price,omitempty"`

	OrderLineItem *OrderLineItem `json:"orderLineItem,omitempty"`

	VersionId string `json:"versionId,omitempty"`

	Quantity float64 `json:"quantity,omitempty"`

	OrderDelivery *OrderDelivery `json:"orderDelivery,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	OrderLineItemId string `json:"orderLineItemId,omitempty"`
}

type OrderDeliveryPositionCollection struct {
	EntityCollection

	Data []OrderDeliveryPosition `json:"data"`
}

package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type OrderDeliveryRepository ClientService

func (t OrderDeliveryRepository) Search(ctx ApiContext, criteria Criteria) (*OrderDeliveryCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/order-delivery", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(OrderDeliveryCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t OrderDeliveryRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/order-delivery", criteria)

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

func (t OrderDeliveryRepository) Upsert(ctx ApiContext, entity []OrderDelivery) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"order_delivery": {
		Entity:  "order_delivery",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t OrderDeliveryRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"order_delivery": {
		Entity:  "order_delivery",
		Action:  "delete",
		Payload: payload,
	}})
}

type OrderDelivery struct {
	ShippingCosts interface{} `json:"shippingCosts,omitempty"`

	ShippingOrderAddress *OrderAddress `json:"shippingOrderAddress,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	OrderVersionId string `json:"orderVersionId,omitempty"`

	ShippingOrderAddressVersionId string `json:"shippingOrderAddressVersionId,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	ShippingDateEarliest time.Time `json:"shippingDateEarliest,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	OrderId string `json:"orderId,omitempty"`

	ShippingMethodId string `json:"shippingMethodId,omitempty"`

	StateMachineState *StateMachineState `json:"stateMachineState,omitempty"`

	TrackingCodes interface{} `json:"trackingCodes,omitempty"`

	ShippingMethod *ShippingMethod `json:"shippingMethod,omitempty"`

	Positions []OrderDeliveryPosition `json:"positions,omitempty"`

	Id string `json:"id,omitempty"`

	VersionId string `json:"versionId,omitempty"`

	ShippingDateLatest time.Time `json:"shippingDateLatest,omitempty"`

	Order *Order `json:"order,omitempty"`

	ShippingOrderAddressId string `json:"shippingOrderAddressId,omitempty"`

	StateId string `json:"stateId,omitempty"`
}

type OrderDeliveryCollection struct {
	EntityCollection

	Data []OrderDelivery `json:"data"`
}

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

func (t OrderDeliveryRepository) SearchAll(ctx ApiContext, criteria Criteria) (*OrderDeliveryCollection, *http.Response, error) {
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
	ShippingMethod *ShippingMethod `json:"shippingMethod,omitempty"`

	Id string `json:"id,omitempty"`

	StateMachineState *StateMachineState `json:"stateMachineState,omitempty"`

	ShippingDateLatest time.Time `json:"shippingDateLatest,omitempty"`

	ShippingCosts interface{} `json:"shippingCosts,omitempty"`

	ShippingOrderAddress *OrderAddress `json:"shippingOrderAddress,omitempty"`

	OrderId string `json:"orderId,omitempty"`

	StateId string `json:"stateId,omitempty"`

	ShippingOrderAddressVersionId string `json:"shippingOrderAddressVersionId,omitempty"`

	ShippingMethodId string `json:"shippingMethodId,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	VersionId string `json:"versionId,omitempty"`

	OrderVersionId string `json:"orderVersionId,omitempty"`

	ShippingDateEarliest time.Time `json:"shippingDateEarliest,omitempty"`

	Order *Order `json:"order,omitempty"`

	Positions []OrderDeliveryPosition `json:"positions,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	ShippingOrderAddressId string `json:"shippingOrderAddressId,omitempty"`

	TrackingCodes interface{} `json:"trackingCodes,omitempty"`
}

type OrderDeliveryCollection struct {
	EntityCollection

	Data []OrderDelivery `json:"data"`
}

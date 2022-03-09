package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type DeliveryTimeRepository ClientService

func (t DeliveryTimeRepository) Search(ctx ApiContext, criteria Criteria) (*DeliveryTimeCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/delivery-time", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(DeliveryTimeCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t DeliveryTimeRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/delivery-time", criteria)

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

func (t DeliveryTimeRepository) Upsert(ctx ApiContext, entity []DeliveryTime) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"delivery_time": {
		Entity:  "delivery_time",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t DeliveryTimeRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"delivery_time": {
		Entity:  "delivery_time",
		Action:  "delete",
		Payload: payload,
	}})
}

type DeliveryTime struct {
	Name string `json:"name,omitempty"`

	Unit string `json:"unit,omitempty"`

	Translations []DeliveryTimeTranslation `json:"translations,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Id string `json:"id,omitempty"`

	Min float64 `json:"min,omitempty"`

	Max float64 `json:"max,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	ShippingMethods []ShippingMethod `json:"shippingMethods,omitempty"`

	Products []Product `json:"products,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Translated interface{} `json:"translated,omitempty"`
}

type DeliveryTimeCollection struct {
	EntityCollection

	Data []DeliveryTime `json:"data"`
}

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

func (t DeliveryTimeRepository) SearchAll(ctx ApiContext, criteria Criteria) (*DeliveryTimeCollection, *http.Response, error) {
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

	Min float64 `json:"min,omitempty"`

	Max float64 `json:"max,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	Translated interface{} `json:"translated,omitempty"`

	Id string `json:"id,omitempty"`

	Unit string `json:"unit,omitempty"`

	ShippingMethods []ShippingMethod `json:"shippingMethods,omitempty"`

	Products []Product `json:"products,omitempty"`

	Translations []DeliveryTimeTranslation `json:"translations,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type DeliveryTimeCollection struct {
	EntityCollection

	Data []DeliveryTime `json:"data"`
}

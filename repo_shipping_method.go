package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type ShippingMethodRepository ClientService

func (t ShippingMethodRepository) Search(ctx ApiContext, criteria Criteria) (*ShippingMethodCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/shipping-method", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(ShippingMethodCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ShippingMethodRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/shipping-method", criteria)

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

func (t ShippingMethodRepository) Upsert(ctx ApiContext, entity []ShippingMethod) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"shipping_method": {
		Entity:  "shipping_method",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t ShippingMethodRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"shipping_method": {
		Entity:  "shipping_method",
		Action:  "delete",
		Payload: payload,
	}})
}

type ShippingMethod struct {
	Name string `json:"name,omitempty"`

	Active bool `json:"active,omitempty"`

	TaxId string `json:"taxId,omitempty"`

	Description string `json:"description,omitempty"`

	TaxType string `json:"taxType,omitempty"`

	TrackingUrl string `json:"trackingUrl,omitempty"`

	AvailabilityRule *Rule `json:"availabilityRule,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	DeliveryTimeId string `json:"deliveryTimeId,omitempty"`

	Translations []ShippingMethodTranslation `json:"translations,omitempty"`

	Prices []ShippingMethodPrice `json:"prices,omitempty"`

	Media *Media `json:"media,omitempty"`

	OrderDeliveries []OrderDelivery `json:"orderDeliveries,omitempty"`

	SalesChannels []SalesChannel `json:"salesChannels,omitempty"`

	SalesChannelDefaultAssignments []SalesChannel `json:"salesChannelDefaultAssignments,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Id string `json:"id,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	AvailabilityRuleId string `json:"availabilityRuleId,omitempty"`

	MediaId string `json:"mediaId,omitempty"`

	DeliveryTime *DeliveryTime `json:"deliveryTime,omitempty"`

	Tags []Tag `json:"tags,omitempty"`

	Tax *Tax `json:"tax,omitempty"`

	Translated interface{} `json:"translated,omitempty"`
}

type ShippingMethodCollection struct {
	EntityCollection

	Data []ShippingMethod `json:"data"`
}

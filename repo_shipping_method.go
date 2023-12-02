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

func (t ShippingMethodRepository) SearchAll(ctx ApiContext, criteria Criteria) (*ShippingMethodCollection, *http.Response, error) {
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

	DeliveryTimeId      string  `json:"deliveryTimeId,omitempty"`

	AvailabilityRule      *Rule  `json:"availabilityRule,omitempty"`

	Prices      []ShippingMethodPrice  `json:"prices,omitempty"`

	SalesChannelDefaultAssignments      []SalesChannel  `json:"salesChannelDefaultAssignments,omitempty"`

	Tax      *Tax  `json:"tax,omitempty"`

	Name      string  `json:"name,omitempty"`

	Position      float64  `json:"position,omitempty"`

	Tags      []Tag  `json:"tags,omitempty"`

	TechnicalName      string  `json:"technicalName,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Active      bool  `json:"active,omitempty"`

	TaxType      string  `json:"taxType,omitempty"`

	Media      *Media  `json:"media,omitempty"`

	OrderDeliveries      []OrderDelivery  `json:"orderDeliveries,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Description      string  `json:"description,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	Id      string  `json:"id,omitempty"`

	AvailabilityRuleId      string  `json:"availabilityRuleId,omitempty"`

	TaxId      string  `json:"taxId,omitempty"`

	DeliveryTime      *DeliveryTime  `json:"deliveryTime,omitempty"`

	AppShippingMethod      *AppShippingMethod  `json:"appShippingMethod,omitempty"`

	TrackingUrl      string  `json:"trackingUrl,omitempty"`

	Translations      []ShippingMethodTranslation  `json:"translations,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	MediaId      string  `json:"mediaId,omitempty"`

	SalesChannels      []SalesChannel  `json:"salesChannels,omitempty"`

}

type ShippingMethodCollection struct {
	EntityCollection

	Data []ShippingMethod `json:"data"`
}

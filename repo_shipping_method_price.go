package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type ShippingMethodPriceRepository ClientService

func (t ShippingMethodPriceRepository) Search(ctx ApiContext, criteria Criteria) (*ShippingMethodPriceCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/shipping-method-price", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(ShippingMethodPriceCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ShippingMethodPriceRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/shipping-method-price", criteria)

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

func (t ShippingMethodPriceRepository) Upsert(ctx ApiContext, entity []ShippingMethodPrice) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"shipping_method_price": {
		Entity:  "shipping_method_price",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t ShippingMethodPriceRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"shipping_method_price": {
		Entity:  "shipping_method_price",
		Action:  "delete",
		Payload: payload,
	}})
}

type ShippingMethodPrice struct {
	Id string `json:"id,omitempty"`

	ShippingMethodId string `json:"shippingMethodId,omitempty"`

	Calculation float64 `json:"calculation,omitempty"`

	CurrencyPrice interface{} `json:"currencyPrice,omitempty"`

	Rule *Rule `json:"rule,omitempty"`

	QuantityStart float64 `json:"quantityStart,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	QuantityEnd float64 `json:"quantityEnd,omitempty"`

	RuleId string `json:"ruleId,omitempty"`

	CalculationRuleId string `json:"calculationRuleId,omitempty"`

	ShippingMethod *ShippingMethod `json:"shippingMethod,omitempty"`

	CalculationRule *Rule `json:"calculationRule,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type ShippingMethodPriceCollection struct {
	EntityCollection

	Data []ShippingMethodPrice `json:"data"`
}

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

func (t ShippingMethodPriceRepository) SearchAll(ctx ApiContext, criteria Criteria) (*ShippingMethodPriceCollection, *http.Response, error) {
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
	ShippingMethodId string `json:"shippingMethodId,omitempty"`

	RuleId string `json:"ruleId,omitempty"`

	Calculation float64 `json:"calculation,omitempty"`

	QuantityStart float64 `json:"quantityStart,omitempty"`

	ShippingMethod *ShippingMethod `json:"shippingMethod,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Id string `json:"id,omitempty"`

	QuantityEnd float64 `json:"quantityEnd,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	CalculationRuleId string `json:"calculationRuleId,omitempty"`

	Rule *Rule `json:"rule,omitempty"`

	CalculationRule *Rule `json:"calculationRule,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	CurrencyPrice interface{} `json:"currencyPrice,omitempty"`
}

type ShippingMethodPriceCollection struct {
	EntityCollection

	Data []ShippingMethodPrice `json:"data"`
}

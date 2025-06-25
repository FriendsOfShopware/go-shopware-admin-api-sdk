package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type ShippingMethodPriceRepository struct {
	*GenericRepository[ShippingMethodPrice]
}

func NewShippingMethodPriceRepository(client *Client) *ShippingMethodPriceRepository {
	return &ShippingMethodPriceRepository{
		GenericRepository: NewGenericRepository[ShippingMethodPrice](client),
	}
}

func (t *ShippingMethodPriceRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[ShippingMethodPrice], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "shipping-method-price")
}

func (t *ShippingMethodPriceRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[ShippingMethodPrice], *http.Response, error) {
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

func (t *ShippingMethodPriceRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "shipping-method-price")
}

func (t *ShippingMethodPriceRepository) Upsert(ctx ApiContext, entity []ShippingMethodPrice) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "shipping_method_price")
}

func (t *ShippingMethodPriceRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "shipping_method_price")
}

type ShippingMethodPrice struct {

	Id      string  `json:"id,omitempty"`

	RuleId      string  `json:"ruleId,omitempty"`

	Calculation      float64  `json:"calculation,omitempty"`

	CalculationRuleId      string  `json:"calculationRuleId,omitempty"`

	QuantityEnd      float64  `json:"quantityEnd,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	ShippingMethod      *ShippingMethod  `json:"shippingMethod,omitempty"`

	Rule      *Rule  `json:"rule,omitempty"`

	ShippingMethodId      string  `json:"shippingMethodId,omitempty"`

	QuantityStart      float64  `json:"quantityStart,omitempty"`

	CurrencyPrice      interface{}  `json:"currencyPrice,omitempty"`

	CalculationRule      *Rule  `json:"calculationRule,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}

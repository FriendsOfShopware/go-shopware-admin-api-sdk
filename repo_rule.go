package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type RuleRepository ClientService

func (t RuleRepository) Search(ctx ApiContext, criteria Criteria) (*RuleCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/rule", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(RuleCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t RuleRepository) SearchAll(ctx ApiContext, criteria Criteria) (*RuleCollection, *http.Response, error) {
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

func (t RuleRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/rule", criteria)

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

func (t RuleRepository) Upsert(ctx ApiContext, entity []Rule) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"rule": {
		Entity:  "rule",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t RuleRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"rule": {
		Entity:  "rule",
		Action:  "delete",
		Payload: payload,
	}})
}

type Rule struct {
	Payload interface{} `json:"payload,omitempty"`

	ShippingMethodPrices []ShippingMethodPrice `json:"shippingMethodPrices,omitempty"`

	PromotionSetGroups []PromotionSetgroup `json:"promotionSetGroups,omitempty"`

	EventActions []EventAction `json:"eventActions,omitempty"`

	Description string `json:"description,omitempty"`

	ModuleTypes interface{} `json:"moduleTypes,omitempty"`

	ProductPrices []ProductPrice `json:"productPrices,omitempty"`

	ShippingMethods []ShippingMethod `json:"shippingMethods,omitempty"`

	PersonaPromotions []Promotion `json:"personaPromotions,omitempty"`

	OrderPromotions []Promotion `json:"orderPromotions,omitempty"`

	Name string `json:"name,omitempty"`

	Priority float64 `json:"priority,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	ShippingMethodPriceCalculations []ShippingMethodPrice `json:"shippingMethodPriceCalculations,omitempty"`

	FlowSequences []FlowSequence `json:"flowSequences,omitempty"`

	CartPromotions []Promotion `json:"cartPromotions,omitempty"`

	PromotionDiscounts []PromotionDiscount `json:"promotionDiscounts,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Id string `json:"id,omitempty"`

	Invalid bool `json:"invalid,omitempty"`

	Conditions []RuleCondition `json:"conditions,omitempty"`

	PaymentMethods []PaymentMethod `json:"paymentMethods,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type RuleCollection struct {
	EntityCollection

	Data []Rule `json:"data"`
}

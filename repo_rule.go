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
	CreatedAt time.Time `json:"createdAt,omitempty"`

	Payload interface{} `json:"payload,omitempty"`

	ProductPrices []ProductPrice `json:"productPrices,omitempty"`

	ShippingMethodPriceCalculations []ShippingMethodPrice `json:"shippingMethodPriceCalculations,omitempty"`

	PersonaPromotions []Promotion `json:"personaPromotions,omitempty"`

	OrderPromotions []Promotion `json:"orderPromotions,omitempty"`

	PromotionDiscounts []PromotionDiscount `json:"promotionDiscounts,omitempty"`

	EventActions []EventAction `json:"eventActions,omitempty"`

	Name string `json:"name,omitempty"`

	Invalid bool `json:"invalid,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	Conditions []RuleCondition `json:"conditions,omitempty"`

	ShippingMethods []ShippingMethod `json:"shippingMethods,omitempty"`

	PaymentMethods []PaymentMethod `json:"paymentMethods,omitempty"`

	Id string `json:"id,omitempty"`

	Priority float64 `json:"priority,omitempty"`

	Description string `json:"description,omitempty"`

	ModuleTypes interface{} `json:"moduleTypes,omitempty"`

	FlowSequences []FlowSequence `json:"flowSequences,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	ShippingMethodPrices []ShippingMethodPrice `json:"shippingMethodPrices,omitempty"`

	CartPromotions []Promotion `json:"cartPromotions,omitempty"`

	PromotionSetGroups []PromotionSetgroup `json:"promotionSetGroups,omitempty"`
}

type RuleCollection struct {
	EntityCollection

	Data []Rule `json:"data"`
}

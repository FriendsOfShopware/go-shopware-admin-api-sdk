package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type PromotionRepository ClientService

func (t PromotionRepository) Search(ctx ApiContext, criteria Criteria) (*PromotionCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/promotion", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(PromotionCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t PromotionRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/promotion", criteria)

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

func (t PromotionRepository) Upsert(ctx ApiContext, entity []Promotion) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"promotion": {
		Entity:  "promotion",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t PromotionRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"promotion": {
		Entity:  "promotion",
		Action:  "delete",
		Payload: payload,
	}})
}

type Promotion struct {
	Discounts []PromotionDiscount `json:"discounts,omitempty"`

	SalesChannels []PromotionSalesChannel `json:"salesChannels,omitempty"`

	Code string `json:"code,omitempty"`

	UseIndividualCodes bool `json:"useIndividualCodes,omitempty"`

	OrdersPerCustomerCount interface{} `json:"ordersPerCustomerCount,omitempty"`

	IndividualCodes []PromotionIndividualCode `json:"individualCodes,omitempty"`

	PersonaCustomers []Customer `json:"personaCustomers,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	Active bool `json:"active,omitempty"`

	PersonaRules []Rule `json:"personaRules,omitempty"`

	OrderRules []Rule `json:"orderRules,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	CustomerRestriction bool `json:"customerRestriction,omitempty"`

	ValidFrom time.Time `json:"validFrom,omitempty"`

	IndividualCodePattern string `json:"individualCodePattern,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Id string `json:"id,omitempty"`

	MaxRedemptionsGlobal float64 `json:"maxRedemptionsGlobal,omitempty"`

	UseCodes bool `json:"useCodes,omitempty"`

	OrderCount float64 `json:"orderCount,omitempty"`

	CartRules []Rule `json:"cartRules,omitempty"`

	Priority float64 `json:"priority,omitempty"`

	ValidUntil time.Time `json:"validUntil,omitempty"`

	MaxRedemptionsPerCustomer float64 `json:"maxRedemptionsPerCustomer,omitempty"`

	Exclusive bool `json:"exclusive,omitempty"`

	UseSetGroups bool `json:"useSetGroups,omitempty"`

	PreventCombination bool `json:"preventCombination,omitempty"`

	Setgroups []PromotionSetgroup `json:"setgroups,omitempty"`

	ExclusionIds interface{} `json:"exclusionIds,omitempty"`

	Name string `json:"name,omitempty"`

	Translated interface{} `json:"translated,omitempty"`

	Translations []PromotionTranslation `json:"translations,omitempty"`
}

type PromotionCollection struct {
	EntityCollection

	Data []Promotion `json:"data"`
}

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

func (t PromotionRepository) SearchAll(ctx ApiContext, criteria Criteria) (*PromotionCollection, *http.Response, error) {
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
	Translations []PromotionTranslation `json:"translations,omitempty"`

	Exclusive bool `json:"exclusive,omitempty"`

	UseIndividualCodes bool `json:"useIndividualCodes,omitempty"`

	SalesChannels []PromotionSalesChannel `json:"salesChannels,omitempty"`

	OrderRules []Rule `json:"orderRules,omitempty"`

	Name string `json:"name,omitempty"`

	MaxRedemptionsGlobal float64 `json:"maxRedemptionsGlobal,omitempty"`

	MaxRedemptionsPerCustomer float64 `json:"maxRedemptionsPerCustomer,omitempty"`

	Discounts []PromotionDiscount `json:"discounts,omitempty"`

	PersonaRules []Rule `json:"personaRules,omitempty"`

	PersonaCustomers []Customer `json:"personaCustomers,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Translated interface{} `json:"translated,omitempty"`

	ValidUntil time.Time `json:"validUntil,omitempty"`

	Priority float64 `json:"priority,omitempty"`

	PreventCombination bool `json:"preventCombination,omitempty"`

	IndividualCodes []PromotionIndividualCode `json:"individualCodes,omitempty"`

	Setgroups []PromotionSetgroup `json:"setgroups,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	OrderCount float64 `json:"orderCount,omitempty"`

	OrdersPerCustomerCount interface{} `json:"ordersPerCustomerCount,omitempty"`

	Active bool `json:"active,omitempty"`

	Code string `json:"code,omitempty"`

	UseSetGroups bool `json:"useSetGroups,omitempty"`

	CustomerRestriction bool `json:"customerRestriction,omitempty"`

	ExclusionIds interface{} `json:"exclusionIds,omitempty"`

	Id string `json:"id,omitempty"`

	IndividualCodePattern string `json:"individualCodePattern,omitempty"`

	ValidFrom time.Time `json:"validFrom,omitempty"`

	UseCodes bool `json:"useCodes,omitempty"`

	CartRules []Rule `json:"cartRules,omitempty"`
}

type PromotionCollection struct {
	EntityCollection

	Data []Promotion `json:"data"`
}

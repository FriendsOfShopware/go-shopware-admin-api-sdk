package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type PromotionRepository struct {
	*GenericRepository[Promotion]
}

func NewPromotionRepository(client *Client) *PromotionRepository {
	return &PromotionRepository{
		GenericRepository: NewGenericRepository[Promotion](client),
	}
}

func (t *PromotionRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[Promotion], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "promotion")
}

func (t *PromotionRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[Promotion], *http.Response, error) {
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

func (t *PromotionRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "promotion")
}

func (t *PromotionRepository) Upsert(ctx ApiContext, entity []Promotion) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "promotion")
}

func (t *PromotionRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "promotion")
}

type Promotion struct {

	Active      bool  `json:"active,omitempty"`

	CartRules      []Rule  `json:"cartRules,omitempty"`

	Code      string  `json:"code,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	CustomerRestriction      bool  `json:"customerRestriction,omitempty"`

	Discounts      []PromotionDiscount  `json:"discounts,omitempty"`

	ExclusionIds      interface{}  `json:"exclusionIds,omitempty"`

	Exclusive      bool  `json:"exclusive,omitempty"`

	Id      string  `json:"id,omitempty"`

	IndividualCodePattern      string  `json:"individualCodePattern,omitempty"`

	IndividualCodes      []PromotionIndividualCode  `json:"individualCodes,omitempty"`

	MaxRedemptionsGlobal      float64  `json:"maxRedemptionsGlobal,omitempty"`

	MaxRedemptionsPerCustomer      float64  `json:"maxRedemptionsPerCustomer,omitempty"`

	Name      string  `json:"name,omitempty"`

	OrderCount      float64  `json:"orderCount,omitempty"`

	OrderLineItems      []OrderLineItem  `json:"orderLineItems,omitempty"`

	OrderRules      []Rule  `json:"orderRules,omitempty"`

	OrdersPerCustomerCount      interface{}  `json:"ordersPerCustomerCount,omitempty"`

	PersonaCustomers      []Customer  `json:"personaCustomers,omitempty"`

	PersonaRules      []Rule  `json:"personaRules,omitempty"`

	PreventCombination      bool  `json:"preventCombination,omitempty"`

	Priority      float64  `json:"priority,omitempty"`

	SalesChannels      []PromotionSalesChannel  `json:"salesChannels,omitempty"`

	Setgroups      []PromotionSetgroup  `json:"setgroups,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	Translations      []PromotionTranslation  `json:"translations,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	UseCodes      bool  `json:"useCodes,omitempty"`

	UseIndividualCodes      bool  `json:"useIndividualCodes,omitempty"`

	UseSetGroups      bool  `json:"useSetGroups,omitempty"`

	ValidFrom      time.Time  `json:"validFrom,omitempty"`

	ValidUntil      time.Time  `json:"validUntil,omitempty"`

}

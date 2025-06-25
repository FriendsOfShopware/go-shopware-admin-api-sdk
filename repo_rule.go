package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type RuleRepository struct {
	*GenericRepository[Rule]
}

func NewRuleRepository(client *Client) *RuleRepository {
	return &RuleRepository{
		GenericRepository: NewGenericRepository[Rule](client),
	}
}

func (t *RuleRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[Rule], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "rule")
}

func (t *RuleRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[Rule], *http.Response, error) {
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

func (t *RuleRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "rule")
}

func (t *RuleRepository) Upsert(ctx ApiContext, entity []Rule) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "rule")
}

func (t *RuleRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "rule")
}

type Rule struct {

	ShippingMethodPriceCalculations      []ShippingMethodPrice  `json:"shippingMethodPriceCalculations,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	ProductPrices      []ProductPrice  `json:"productPrices,omitempty"`

	Name      string  `json:"name,omitempty"`

	Priority      float64  `json:"priority,omitempty"`

	Description      string  `json:"description,omitempty"`

	Payload      interface{}  `json:"payload,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	ModuleTypes      interface{}  `json:"moduleTypes,omitempty"`

	Conditions      []RuleCondition  `json:"conditions,omitempty"`

	ShippingMethods      []ShippingMethod  `json:"shippingMethods,omitempty"`

	PaymentMethods      []PaymentMethod  `json:"paymentMethods,omitempty"`

	PersonaPromotions      []Promotion  `json:"personaPromotions,omitempty"`

	TaxProviders      []TaxProvider  `json:"taxProviders,omitempty"`

	Tags      []Tag  `json:"tags,omitempty"`

	OrderPromotions      []Promotion  `json:"orderPromotions,omitempty"`

	PromotionSetGroups      []PromotionSetgroup  `json:"promotionSetGroups,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	ShippingMethodPrices      []ShippingMethodPrice  `json:"shippingMethodPrices,omitempty"`

	Areas      interface{}  `json:"areas,omitempty"`

	FlowSequences      []FlowSequence  `json:"flowSequences,omitempty"`

	CartPromotions      []Promotion  `json:"cartPromotions,omitempty"`

	PromotionDiscounts      []PromotionDiscount  `json:"promotionDiscounts,omitempty"`

	Id      string  `json:"id,omitempty"`

	Invalid      bool  `json:"invalid,omitempty"`

}

package go_shopware_admin_sdk

import (
	"net/http"

)

type PromotionCartRuleRepository struct {
	*GenericRepository[PromotionCartRule]
}

func NewPromotionCartRuleRepository(client *Client) *PromotionCartRuleRepository {
	return &PromotionCartRuleRepository{
		GenericRepository: NewGenericRepository[PromotionCartRule](client),
	}
}

func (t *PromotionCartRuleRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[PromotionCartRule], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "promotion-cart-rule")
}

func (t *PromotionCartRuleRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[PromotionCartRule], *http.Response, error) {
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

func (t *PromotionCartRuleRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "promotion-cart-rule")
}

func (t *PromotionCartRuleRepository) Upsert(ctx ApiContext, entity []PromotionCartRule) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "promotion_cart_rule")
}

func (t *PromotionCartRuleRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "promotion_cart_rule")
}

type PromotionCartRule struct {

	Promotion      *Promotion  `json:"promotion,omitempty"`

	PromotionId      string  `json:"promotionId,omitempty"`

	Rule      *Rule  `json:"rule,omitempty"`

	RuleId      string  `json:"ruleId,omitempty"`

}

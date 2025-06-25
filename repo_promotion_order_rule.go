package go_shopware_admin_sdk

import (
	"net/http"

)

type PromotionOrderRuleRepository struct {
	*GenericRepository[PromotionOrderRule]
}

func NewPromotionOrderRuleRepository(client *Client) *PromotionOrderRuleRepository {
	return &PromotionOrderRuleRepository{
		GenericRepository: NewGenericRepository[PromotionOrderRule](client),
	}
}

func (t *PromotionOrderRuleRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[PromotionOrderRule], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "promotion-order-rule")
}

func (t *PromotionOrderRuleRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[PromotionOrderRule], *http.Response, error) {
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

func (t *PromotionOrderRuleRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "promotion-order-rule")
}

func (t *PromotionOrderRuleRepository) Upsert(ctx ApiContext, entity []PromotionOrderRule) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "promotion_order_rule")
}

func (t *PromotionOrderRuleRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "promotion_order_rule")
}

type PromotionOrderRule struct {

	Promotion      *Promotion  `json:"promotion,omitempty"`

	PromotionId      string  `json:"promotionId,omitempty"`

	Rule      *Rule  `json:"rule,omitempty"`

	RuleId      string  `json:"ruleId,omitempty"`

}

package go_shopware_admin_sdk

import (
	"net/http"

)

type PromotionDiscountRuleRepository struct {
	*GenericRepository[PromotionDiscountRule]
}

func NewPromotionDiscountRuleRepository(client *Client) *PromotionDiscountRuleRepository {
	return &PromotionDiscountRuleRepository{
		GenericRepository: NewGenericRepository[PromotionDiscountRule](client),
	}
}

func (t *PromotionDiscountRuleRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[PromotionDiscountRule], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "promotion-discount-rule")
}

func (t *PromotionDiscountRuleRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[PromotionDiscountRule], *http.Response, error) {
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

func (t *PromotionDiscountRuleRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "promotion-discount-rule")
}

func (t *PromotionDiscountRuleRepository) Upsert(ctx ApiContext, entity []PromotionDiscountRule) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "promotion_discount_rule")
}

func (t *PromotionDiscountRuleRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "promotion_discount_rule")
}

type PromotionDiscountRule struct {

	DiscountId      string  `json:"discountId,omitempty"`

	RuleId      string  `json:"ruleId,omitempty"`

	Discount      *PromotionDiscount  `json:"discount,omitempty"`

	Rule      *Rule  `json:"rule,omitempty"`

}

package go_shopware_admin_sdk

import (
	"net/http"

)

type PromotionSetgroupRuleRepository struct {
	*GenericRepository[PromotionSetgroupRule]
}

func NewPromotionSetgroupRuleRepository(client *Client) *PromotionSetgroupRuleRepository {
	return &PromotionSetgroupRuleRepository{
		GenericRepository: NewGenericRepository[PromotionSetgroupRule](client),
	}
}

func (t *PromotionSetgroupRuleRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[PromotionSetgroupRule], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "promotion-setgroup-rule")
}

func (t *PromotionSetgroupRuleRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[PromotionSetgroupRule], *http.Response, error) {
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

func (t *PromotionSetgroupRuleRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "promotion-setgroup-rule")
}

func (t *PromotionSetgroupRuleRepository) Upsert(ctx ApiContext, entity []PromotionSetgroupRule) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "promotion_setgroup_rule")
}

func (t *PromotionSetgroupRuleRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "promotion_setgroup_rule")
}

type PromotionSetgroupRule struct {

	SetgroupId      string  `json:"setgroupId,omitempty"`

	RuleId      string  `json:"ruleId,omitempty"`

	Setgroup      *PromotionSetgroup  `json:"setgroup,omitempty"`

	Rule      *Rule  `json:"rule,omitempty"`

}

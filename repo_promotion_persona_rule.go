package go_shopware_admin_sdk

import (
	"net/http"

)

type PromotionPersonaRuleRepository struct {
	*GenericRepository[PromotionPersonaRule]
}

func NewPromotionPersonaRuleRepository(client *Client) *PromotionPersonaRuleRepository {
	return &PromotionPersonaRuleRepository{
		GenericRepository: NewGenericRepository[PromotionPersonaRule](client),
	}
}

func (t *PromotionPersonaRuleRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[PromotionPersonaRule], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "promotion-persona-rule")
}

func (t *PromotionPersonaRuleRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[PromotionPersonaRule], *http.Response, error) {
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

func (t *PromotionPersonaRuleRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "promotion-persona-rule")
}

func (t *PromotionPersonaRuleRepository) Upsert(ctx ApiContext, entity []PromotionPersonaRule) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "promotion_persona_rule")
}

func (t *PromotionPersonaRuleRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "promotion_persona_rule")
}

type PromotionPersonaRule struct {

	PromotionId      string  `json:"promotionId,omitempty"`

	RuleId      string  `json:"ruleId,omitempty"`

	Promotion      *Promotion  `json:"promotion,omitempty"`

	Rule      *Rule  `json:"rule,omitempty"`

}

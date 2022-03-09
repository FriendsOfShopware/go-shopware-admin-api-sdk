package go_shopware_admin_sdk

import (
	"net/http"
)

type PromotionPersonaRuleRepository ClientService

func (t PromotionPersonaRuleRepository) Search(ctx ApiContext, criteria Criteria) (*PromotionPersonaRuleCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/promotion-persona-rule", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(PromotionPersonaRuleCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t PromotionPersonaRuleRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/promotion-persona-rule", criteria)

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

func (t PromotionPersonaRuleRepository) Upsert(ctx ApiContext, entity []PromotionPersonaRule) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"promotion_persona_rule": {
		Entity:  "promotion_persona_rule",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t PromotionPersonaRuleRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"promotion_persona_rule": {
		Entity:  "promotion_persona_rule",
		Action:  "delete",
		Payload: payload,
	}})
}

type PromotionPersonaRule struct {
	PromotionId string `json:"promotionId,omitempty"`

	RuleId string `json:"ruleId,omitempty"`

	Promotion *Promotion `json:"promotion,omitempty"`

	Rule *Rule `json:"rule,omitempty"`
}

type PromotionPersonaRuleCollection struct {
	EntityCollection

	Data []PromotionPersonaRule `json:"data"`
}

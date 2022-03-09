package go_shopware_admin_sdk

import (
	"net/http"
)

type PromotionSetgroupRuleRepository ClientService

func (t PromotionSetgroupRuleRepository) Search(ctx ApiContext, criteria Criteria) (*PromotionSetgroupRuleCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/promotion-setgroup-rule", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(PromotionSetgroupRuleCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t PromotionSetgroupRuleRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/promotion-setgroup-rule", criteria)

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

func (t PromotionSetgroupRuleRepository) Upsert(ctx ApiContext, entity []PromotionSetgroupRule) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"promotion_setgroup_rule": {
		Entity:  "promotion_setgroup_rule",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t PromotionSetgroupRuleRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"promotion_setgroup_rule": {
		Entity:  "promotion_setgroup_rule",
		Action:  "delete",
		Payload: payload,
	}})
}

type PromotionSetgroupRule struct {
	Setgroup *PromotionSetgroup `json:"setgroup,omitempty"`

	Rule *Rule `json:"rule,omitempty"`

	SetgroupId string `json:"setgroupId,omitempty"`

	RuleId string `json:"ruleId,omitempty"`
}

type PromotionSetgroupRuleCollection struct {
	EntityCollection

	Data []PromotionSetgroupRule `json:"data"`
}

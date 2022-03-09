package go_shopware_admin_sdk

import (
	"net/http"
)

type PromotionOrderRuleRepository ClientService

func (t PromotionOrderRuleRepository) Search(ctx ApiContext, criteria Criteria) (*PromotionOrderRuleCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/promotion-order-rule", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(PromotionOrderRuleCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t PromotionOrderRuleRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/promotion-order-rule", criteria)

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

func (t PromotionOrderRuleRepository) Upsert(ctx ApiContext, entity []PromotionOrderRule) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"promotion_order_rule": {
		Entity:  "promotion_order_rule",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t PromotionOrderRuleRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"promotion_order_rule": {
		Entity:  "promotion_order_rule",
		Action:  "delete",
		Payload: payload,
	}})
}

type PromotionOrderRule struct {
	RuleId string `json:"ruleId,omitempty"`

	Promotion *Promotion `json:"promotion,omitempty"`

	Rule *Rule `json:"rule,omitempty"`

	PromotionId string `json:"promotionId,omitempty"`
}

type PromotionOrderRuleCollection struct {
	EntityCollection

	Data []PromotionOrderRule `json:"data"`
}

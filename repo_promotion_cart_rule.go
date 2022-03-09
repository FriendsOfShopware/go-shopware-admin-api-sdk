package go_shopware_admin_sdk

import (
	"net/http"
)

type PromotionCartRuleRepository ClientService

func (t PromotionCartRuleRepository) Search(ctx ApiContext, criteria Criteria) (*PromotionCartRuleCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/promotion-cart-rule", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(PromotionCartRuleCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t PromotionCartRuleRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/promotion-cart-rule", criteria)

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

func (t PromotionCartRuleRepository) Upsert(ctx ApiContext, entity []PromotionCartRule) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"promotion_cart_rule": {
		Entity:  "promotion_cart_rule",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t PromotionCartRuleRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"promotion_cart_rule": {
		Entity:  "promotion_cart_rule",
		Action:  "delete",
		Payload: payload,
	}})
}

type PromotionCartRule struct {
	PromotionId string `json:"promotionId,omitempty"`

	RuleId string `json:"ruleId,omitempty"`

	Promotion *Promotion `json:"promotion,omitempty"`

	Rule *Rule `json:"rule,omitempty"`
}

type PromotionCartRuleCollection struct {
	EntityCollection

	Data []PromotionCartRule `json:"data"`
}

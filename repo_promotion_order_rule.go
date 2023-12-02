package go_shopware_admin_sdk

import (
	"net/http"
	"time"
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

func (t PromotionOrderRuleRepository) SearchAll(ctx ApiContext, criteria Criteria) (*PromotionOrderRuleCollection, *http.Response, error) {
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

	PromotionId      string  `json:"promotionId,omitempty"`

	RuleId      string  `json:"ruleId,omitempty"`

	Promotion      *Promotion  `json:"promotion,omitempty"`

	Rule      *Rule  `json:"rule,omitempty"`

}

type PromotionOrderRuleCollection struct {
	EntityCollection

	Data []PromotionOrderRule `json:"data"`
}

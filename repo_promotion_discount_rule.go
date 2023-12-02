package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type PromotionDiscountRuleRepository ClientService

func (t PromotionDiscountRuleRepository) Search(ctx ApiContext, criteria Criteria) (*PromotionDiscountRuleCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/promotion-discount-rule", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(PromotionDiscountRuleCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t PromotionDiscountRuleRepository) SearchAll(ctx ApiContext, criteria Criteria) (*PromotionDiscountRuleCollection, *http.Response, error) {
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

func (t PromotionDiscountRuleRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/promotion-discount-rule", criteria)

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

func (t PromotionDiscountRuleRepository) Upsert(ctx ApiContext, entity []PromotionDiscountRule) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"promotion_discount_rule": {
		Entity:  "promotion_discount_rule",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t PromotionDiscountRuleRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"promotion_discount_rule": {
		Entity:  "promotion_discount_rule",
		Action:  "delete",
		Payload: payload,
	}})
}

type PromotionDiscountRule struct {

	DiscountId      string  `json:"discountId,omitempty"`

	RuleId      string  `json:"ruleId,omitempty"`

	Discount      *PromotionDiscount  `json:"discount,omitempty"`

	Rule      *Rule  `json:"rule,omitempty"`

}

type PromotionDiscountRuleCollection struct {
	EntityCollection

	Data []PromotionDiscountRule `json:"data"`
}

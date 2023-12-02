package go_shopware_admin_sdk

import (
	"net/http"

	"time"
)

type PromotionDiscountRepository ClientService

func (t PromotionDiscountRepository) Search(ctx ApiContext, criteria Criteria) (*PromotionDiscountCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/promotion-discount", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(PromotionDiscountCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t PromotionDiscountRepository) SearchAll(ctx ApiContext, criteria Criteria) (*PromotionDiscountCollection, *http.Response, error) {
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

func (t PromotionDiscountRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/promotion-discount", criteria)

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

func (t PromotionDiscountRepository) Upsert(ctx ApiContext, entity []PromotionDiscount) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"promotion_discount": {
		Entity:  "promotion_discount",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t PromotionDiscountRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"promotion_discount": {
		Entity:  "promotion_discount",
		Action:  "delete",
		Payload: payload,
	}})
}

type PromotionDiscount struct {
	SorterKey string `json:"sorterKey,omitempty"`

	PromotionDiscountPrices []PromotionDiscountPrices `json:"promotionDiscountPrices,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	PromotionId string `json:"promotionId,omitempty"`

	Scope string `json:"scope,omitempty"`

	Type string `json:"type,omitempty"`

	MaxValue float64 `json:"maxValue,omitempty"`

	ApplierKey string `json:"applierKey,omitempty"`

	Promotion *Promotion `json:"promotion,omitempty"`

	DiscountRules []Rule `json:"discountRules,omitempty"`

	Id string `json:"id,omitempty"`

	Value float64 `json:"value,omitempty"`

	ConsiderAdvancedRules bool `json:"considerAdvancedRules,omitempty"`

	PickerKey string `json:"pickerKey,omitempty"`

	UsageKey string `json:"usageKey,omitempty"`
}

type PromotionDiscountCollection struct {
	EntityCollection

	Data []PromotionDiscount `json:"data"`
}

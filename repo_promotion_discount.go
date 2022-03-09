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
	PromotionDiscountPrices []PromotionDiscountPrices `json:"promotionDiscountPrices,omitempty"`

	DiscountRules []Rule `json:"discountRules,omitempty"`

	ConsiderAdvancedRules bool `json:"considerAdvancedRules,omitempty"`

	SorterKey string `json:"sorterKey,omitempty"`

	ApplierKey string `json:"applierKey,omitempty"`

	Promotion *Promotion `json:"promotion,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	PromotionId string `json:"promotionId,omitempty"`

	Type string `json:"type,omitempty"`

	Value float64 `json:"value,omitempty"`

	MaxValue float64 `json:"maxValue,omitempty"`

	UsageKey string `json:"usageKey,omitempty"`

	Id string `json:"id,omitempty"`

	PickerKey string `json:"pickerKey,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Scope string `json:"scope,omitempty"`
}

type PromotionDiscountCollection struct {
	EntityCollection

	Data []PromotionDiscount `json:"data"`
}

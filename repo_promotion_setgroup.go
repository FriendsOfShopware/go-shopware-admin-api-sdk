package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type PromotionSetgroupRepository ClientService

func (t PromotionSetgroupRepository) Search(ctx ApiContext, criteria Criteria) (*PromotionSetgroupCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/promotion-setgroup", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(PromotionSetgroupCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t PromotionSetgroupRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/promotion-setgroup", criteria)

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

func (t PromotionSetgroupRepository) Upsert(ctx ApiContext, entity []PromotionSetgroup) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"promotion_setgroup": {
		Entity:  "promotion_setgroup",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t PromotionSetgroupRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"promotion_setgroup": {
		Entity:  "promotion_setgroup",
		Action:  "delete",
		Payload: payload,
	}})
}

type PromotionSetgroup struct {
	SetGroupRules []Rule `json:"setGroupRules,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	PromotionId string `json:"promotionId,omitempty"`

	PackagerKey string `json:"packagerKey,omitempty"`

	SorterKey string `json:"sorterKey,omitempty"`

	Promotion *Promotion `json:"promotion,omitempty"`

	Id string `json:"id,omitempty"`

	Value float64 `json:"value,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`
}

type PromotionSetgroupCollection struct {
	EntityCollection

	Data []PromotionSetgroup `json:"data"`
}

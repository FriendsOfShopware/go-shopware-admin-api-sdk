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

func (t PromotionSetgroupRepository) SearchAll(ctx ApiContext, criteria Criteria) (*PromotionSetgroupCollection, *http.Response, error) {
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
	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	PackagerKey string `json:"packagerKey,omitempty"`

	SorterKey string `json:"sorterKey,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Promotion *Promotion `json:"promotion,omitempty"`

	SetGroupRules []Rule `json:"setGroupRules,omitempty"`

	Id string `json:"id,omitempty"`

	PromotionId string `json:"promotionId,omitempty"`

	Value float64 `json:"value,omitempty"`
}

type PromotionSetgroupCollection struct {
	EntityCollection

	Data []PromotionSetgroup `json:"data"`
}

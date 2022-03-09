package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type PromotionSalesChannelRepository ClientService

func (t PromotionSalesChannelRepository) Search(ctx ApiContext, criteria Criteria) (*PromotionSalesChannelCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/promotion-sales-channel", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(PromotionSalesChannelCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t PromotionSalesChannelRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/promotion-sales-channel", criteria)

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

func (t PromotionSalesChannelRepository) Upsert(ctx ApiContext, entity []PromotionSalesChannel) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"promotion_sales_channel": {
		Entity:  "promotion_sales_channel",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t PromotionSalesChannelRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"promotion_sales_channel": {
		Entity:  "promotion_sales_channel",
		Action:  "delete",
		Payload: payload,
	}})
}

type PromotionSalesChannel struct {
	Priority float64 `json:"priority,omitempty"`

	Promotion *Promotion `json:"promotion,omitempty"`

	SalesChannel *SalesChannel `json:"salesChannel,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Id string `json:"id,omitempty"`

	PromotionId string `json:"promotionId,omitempty"`

	SalesChannelId string `json:"salesChannelId,omitempty"`
}

type PromotionSalesChannelCollection struct {
	EntityCollection

	Data []PromotionSalesChannel `json:"data"`
}

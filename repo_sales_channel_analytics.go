package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type SalesChannelAnalyticsRepository ClientService

func (t SalesChannelAnalyticsRepository) Search(ctx ApiContext, criteria Criteria) (*SalesChannelAnalyticsCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/sales-channel-analytics", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(SalesChannelAnalyticsCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t SalesChannelAnalyticsRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/sales-channel-analytics", criteria)

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

func (t SalesChannelAnalyticsRepository) Upsert(ctx ApiContext, entity []SalesChannelAnalytics) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"sales_channel_analytics": {
		Entity:  "sales_channel_analytics",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t SalesChannelAnalyticsRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"sales_channel_analytics": {
		Entity:  "sales_channel_analytics",
		Action:  "delete",
		Payload: payload,
	}})
}

type SalesChannelAnalytics struct {
	SalesChannel *SalesChannel `json:"salesChannel,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Id string `json:"id,omitempty"`

	TrackingId string `json:"trackingId,omitempty"`

	Active bool `json:"active,omitempty"`

	TrackOrders bool `json:"trackOrders,omitempty"`

	AnonymizeIp bool `json:"anonymizeIp,omitempty"`
}

type SalesChannelAnalyticsCollection struct {
	EntityCollection

	Data []SalesChannelAnalytics `json:"data"`
}

package go_shopware_admin_sdk

import (
	"net/http"
)

type EventActionSalesChannelRepository ClientService

func (t EventActionSalesChannelRepository) Search(ctx ApiContext, criteria Criteria) (*EventActionSalesChannelCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/event-action-sales-channel", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(EventActionSalesChannelCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t EventActionSalesChannelRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/event-action-sales-channel", criteria)

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

func (t EventActionSalesChannelRepository) Upsert(ctx ApiContext, entity []EventActionSalesChannel) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"event_action_sales_channel": {
		Entity:  "event_action_sales_channel",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t EventActionSalesChannelRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"event_action_sales_channel": {
		Entity:  "event_action_sales_channel",
		Action:  "delete",
		Payload: payload,
	}})
}

type EventActionSalesChannel struct {
	EventActionId string `json:"eventActionId,omitempty"`

	SalesChannelId string `json:"salesChannelId,omitempty"`

	EventAction *EventAction `json:"eventAction,omitempty"`

	SalesChannel *SalesChannel `json:"salesChannel,omitempty"`
}

type EventActionSalesChannelCollection struct {
	EntityCollection

	Data []EventActionSalesChannel `json:"data"`
}

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

func (t EventActionSalesChannelRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EventActionSalesChannelCollection, *http.Response, error) {
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
	SalesChannel *SalesChannel `json:"salesChannel,omitempty"`

	EventActionId string `json:"eventActionId,omitempty"`

	SalesChannelId string `json:"salesChannelId,omitempty"`

	EventAction *EventAction `json:"eventAction,omitempty"`
}

type EventActionSalesChannelCollection struct {
	EntityCollection

	Data []EventActionSalesChannel `json:"data"`
}

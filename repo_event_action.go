package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type EventActionRepository ClientService

func (t EventActionRepository) Search(ctx ApiContext, criteria Criteria) (*EventActionCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/event-action", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(EventActionCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t EventActionRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EventActionCollection, *http.Response, error) {
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

func (t EventActionRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/event-action", criteria)

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

func (t EventActionRepository) Upsert(ctx ApiContext, entity []EventAction) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"event_action": {
		Entity:  "event_action",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t EventActionRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"event_action": {
		Entity:  "event_action",
		Action:  "delete",
		Payload: payload,
	}})
}

type EventAction struct {
	Active bool `json:"active,omitempty"`

	Title string `json:"title,omitempty"`

	Rules []Rule `json:"rules,omitempty"`

	SalesChannels []SalesChannel `json:"salesChannels,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	EventName string `json:"eventName,omitempty"`

	ActionName string `json:"actionName,omitempty"`

	Config interface{} `json:"config,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Id string `json:"id,omitempty"`
}

type EventActionCollection struct {
	EntityCollection

	Data []EventAction `json:"data"`
}

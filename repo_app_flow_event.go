package go_shopware_admin_sdk

import (
	"net/http"

	"time"
)

type AppFlowEventRepository ClientService

func (t AppFlowEventRepository) Search(ctx ApiContext, criteria Criteria) (*AppFlowEventCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/app-flow-event", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(AppFlowEventCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t AppFlowEventRepository) SearchAll(ctx ApiContext, criteria Criteria) (*AppFlowEventCollection, *http.Response, error) {
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

func (t AppFlowEventRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/app-flow-event", criteria)

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

func (t AppFlowEventRepository) Upsert(ctx ApiContext, entity []AppFlowEvent) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"app_flow_event": {
		Entity:  "app_flow_event",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t AppFlowEventRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"app_flow_event": {
		Entity:  "app_flow_event",
		Action:  "delete",
		Payload: payload,
	}})
}

type AppFlowEvent struct {
	Name string `json:"name,omitempty"`

	Aware interface{} `json:"aware,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Id string `json:"id,omitempty"`

	AppId string `json:"appId,omitempty"`

	Flows []Flow `json:"flows,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	App *App `json:"app,omitempty"`
}

type AppFlowEventCollection struct {
	EntityCollection

	Data []AppFlowEvent `json:"data"`
}

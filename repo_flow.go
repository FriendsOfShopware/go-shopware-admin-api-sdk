package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type FlowRepository ClientService

func (t FlowRepository) Search(ctx ApiContext, criteria Criteria) (*FlowCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/flow", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(FlowCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t FlowRepository) SearchAll(ctx ApiContext, criteria Criteria) (*FlowCollection, *http.Response, error) {
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

func (t FlowRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/flow", criteria)

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

func (t FlowRepository) Upsert(ctx ApiContext, entity []Flow) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"flow": {
		Entity:  "flow",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t FlowRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"flow": {
		Entity:  "flow",
		Action:  "delete",
		Payload: payload,
	}})
}

type Flow struct {
	EventName string `json:"eventName,omitempty"`

	Payload interface{} `json:"payload,omitempty"`

	Active bool `json:"active,omitempty"`

	Description string `json:"description,omitempty"`

	Sequences []FlowSequence `json:"sequences,omitempty"`

	Id string `json:"id,omitempty"`

	Priority float64 `json:"priority,omitempty"`

	Invalid bool `json:"invalid,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Name string `json:"name,omitempty"`
}

type FlowCollection struct {
	EntityCollection

	Data []Flow `json:"data"`
}

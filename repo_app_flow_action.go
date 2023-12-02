package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type AppFlowActionRepository ClientService

func (t AppFlowActionRepository) Search(ctx ApiContext, criteria Criteria) (*AppFlowActionCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/app-flow-action", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(AppFlowActionCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t AppFlowActionRepository) SearchAll(ctx ApiContext, criteria Criteria) (*AppFlowActionCollection, *http.Response, error) {
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

func (t AppFlowActionRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/app-flow-action", criteria)

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

func (t AppFlowActionRepository) Upsert(ctx ApiContext, entity []AppFlowAction) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"app_flow_action": {
		Entity:  "app_flow_action",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t AppFlowActionRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"app_flow_action": {
		Entity:  "app_flow_action",
		Action:  "delete",
		Payload: payload,
	}})
}

type AppFlowAction struct {

	App      *App  `json:"app,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	Badge      string  `json:"badge,omitempty"`

	Parameters      interface{}  `json:"parameters,omitempty"`

	Headers      interface{}  `json:"headers,omitempty"`

	IconRaw      interface{}  `json:"iconRaw,omitempty"`

	Headline      string  `json:"headline,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	FlowSequences      []FlowSequence  `json:"flowSequences,omitempty"`

	Id      string  `json:"id,omitempty"`

	Name      string  `json:"name,omitempty"`

	Config      interface{}  `json:"config,omitempty"`

	Url      string  `json:"url,omitempty"`

	Label      string  `json:"label,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	AppId      string  `json:"appId,omitempty"`

	Icon      string  `json:"icon,omitempty"`

	SwIcon      string  `json:"swIcon,omitempty"`

	Delayable      bool  `json:"delayable,omitempty"`

	Description      string  `json:"description,omitempty"`

	Requirements      interface{}  `json:"requirements,omitempty"`

	Translations      []AppFlowActionTranslation  `json:"translations,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

}

type AppFlowActionCollection struct {
	EntityCollection

	Data []AppFlowAction `json:"data"`
}

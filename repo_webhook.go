package go_shopware_admin_sdk

import (
	"net/http"

	"time"
)

type WebhookRepository ClientService

func (t WebhookRepository) Search(ctx ApiContext, criteria Criteria) (*WebhookCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/webhook", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(WebhookCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t WebhookRepository) SearchAll(ctx ApiContext, criteria Criteria) (*WebhookCollection, *http.Response, error) {
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

func (t WebhookRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/webhook", criteria)

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

func (t WebhookRepository) Upsert(ctx ApiContext, entity []Webhook) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"webhook": {
		Entity:  "webhook",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t WebhookRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"webhook": {
		Entity:  "webhook",
		Action:  "delete",
		Payload: payload,
	}})
}

type Webhook struct {
	Url string `json:"url,omitempty"`

	ErrorCount float64 `json:"errorCount,omitempty"`

	AppId string `json:"appId,omitempty"`

	App *App `json:"app,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	EventName string `json:"eventName,omitempty"`

	OnlyLiveVersion bool `json:"onlyLiveVersion,omitempty"`

	Active bool `json:"active,omitempty"`
}

type WebhookCollection struct {
	EntityCollection

	Data []Webhook `json:"data"`
}

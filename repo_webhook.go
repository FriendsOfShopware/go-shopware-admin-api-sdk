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
	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Id string `json:"id,omitempty"`

	Url string `json:"url,omitempty"`

	Active bool `json:"active,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	App *App `json:"app,omitempty"`

	Name string `json:"name,omitempty"`

	EventName string `json:"eventName,omitempty"`

	ErrorCount float64 `json:"errorCount,omitempty"`

	AppId string `json:"appId,omitempty"`
}

type WebhookCollection struct {
	EntityCollection

	Data []Webhook `json:"data"`
}

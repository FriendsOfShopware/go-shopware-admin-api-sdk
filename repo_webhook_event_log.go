package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type WebhookEventLogRepository ClientService

func (t WebhookEventLogRepository) Search(ctx ApiContext, criteria Criteria) (*WebhookEventLogCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/webhook-event-log", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(WebhookEventLogCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t WebhookEventLogRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/webhook-event-log", criteria)

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

func (t WebhookEventLogRepository) Upsert(ctx ApiContext, entity []WebhookEventLog) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"webhook_event_log": {
		Entity:  "webhook_event_log",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t WebhookEventLogRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"webhook_event_log": {
		Entity:  "webhook_event_log",
		Action:  "delete",
		Payload: payload,
	}})
}

type WebhookEventLog struct {
	WebhookName string `json:"webhookName,omitempty"`

	Timestamp float64 `json:"timestamp,omitempty"`

	ProcessingTime float64 `json:"processingTime,omitempty"`

	AppVersion string `json:"appVersion,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Id string `json:"id,omitempty"`

	ResponseReasonPhrase string `json:"responseReasonPhrase,omitempty"`

	RequestContent interface{} `json:"requestContent,omitempty"`

	ResponseContent interface{} `json:"responseContent,omitempty"`

	ResponseStatusCode float64 `json:"responseStatusCode,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	AppName string `json:"appName,omitempty"`

	EventName string `json:"eventName,omitempty"`

	DeliveryStatus string `json:"deliveryStatus,omitempty"`

	Url string `json:"url,omitempty"`

	SerializedWebhookMessage interface{} `json:"serializedWebhookMessage,omitempty"`
}

type WebhookEventLogCollection struct {
	EntityCollection

	Data []WebhookEventLog `json:"data"`
}

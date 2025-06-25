package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type WebhookEventLogRepository struct {
	*GenericRepository[WebhookEventLog]
}

func NewWebhookEventLogRepository(client *Client) *WebhookEventLogRepository {
	return &WebhookEventLogRepository{
		GenericRepository: NewGenericRepository[WebhookEventLog](client),
	}
}

func (t *WebhookEventLogRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[WebhookEventLog], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "webhook-event-log")
}

func (t *WebhookEventLogRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[WebhookEventLog], *http.Response, error) {
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

func (t *WebhookEventLogRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "webhook-event-log")
}

func (t *WebhookEventLogRepository) Upsert(ctx ApiContext, entity []WebhookEventLog) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "webhook_event_log")
}

func (t *WebhookEventLogRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "webhook_event_log")
}

type WebhookEventLog struct {

	AppName      string  `json:"appName,omitempty"`

	AppVersion      string  `json:"appVersion,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	DeliveryStatus      string  `json:"deliveryStatus,omitempty"`

	EventName      string  `json:"eventName,omitempty"`

	Id      string  `json:"id,omitempty"`

	OnlyLiveVersion      bool  `json:"onlyLiveVersion,omitempty"`

	ProcessingTime      float64  `json:"processingTime,omitempty"`

	RequestContent      interface{}  `json:"requestContent,omitempty"`

	ResponseContent      interface{}  `json:"responseContent,omitempty"`

	ResponseReasonPhrase      string  `json:"responseReasonPhrase,omitempty"`

	ResponseStatusCode      float64  `json:"responseStatusCode,omitempty"`

	SerializedWebhookMessage      interface{}  `json:"serializedWebhookMessage,omitempty"`

	Timestamp      float64  `json:"timestamp,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Url      string  `json:"url,omitempty"`

	WebhookName      string  `json:"webhookName,omitempty"`

}

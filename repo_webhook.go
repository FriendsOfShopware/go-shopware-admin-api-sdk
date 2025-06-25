package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type WebhookRepository struct {
	*GenericRepository[Webhook]
}

func NewWebhookRepository(client *Client) *WebhookRepository {
	return &WebhookRepository{
		GenericRepository: NewGenericRepository[Webhook](client),
	}
}

func (t *WebhookRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[Webhook], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "webhook")
}

func (t *WebhookRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[Webhook], *http.Response, error) {
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

func (t *WebhookRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "webhook")
}

func (t *WebhookRepository) Upsert(ctx ApiContext, entity []Webhook) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "webhook")
}

func (t *WebhookRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "webhook")
}

type Webhook struct {

	AppId      string  `json:"appId,omitempty"`

	App      *App  `json:"app,omitempty"`

	EventName      string  `json:"eventName,omitempty"`

	Url      string  `json:"url,omitempty"`

	Active      bool  `json:"active,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Id      string  `json:"id,omitempty"`

	Name      string  `json:"name,omitempty"`

	OnlyLiveVersion      bool  `json:"onlyLiveVersion,omitempty"`

	ErrorCount      float64  `json:"errorCount,omitempty"`

}

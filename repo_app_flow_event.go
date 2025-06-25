package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type AppFlowEventRepository struct {
	*GenericRepository[AppFlowEvent]
}

func NewAppFlowEventRepository(client *Client) *AppFlowEventRepository {
	return &AppFlowEventRepository{
		GenericRepository: NewGenericRepository[AppFlowEvent](client),
	}
}

func (t *AppFlowEventRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[AppFlowEvent], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "app-flow-event")
}

func (t *AppFlowEventRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[AppFlowEvent], *http.Response, error) {
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

func (t *AppFlowEventRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "app-flow-event")
}

func (t *AppFlowEventRepository) Upsert(ctx ApiContext, entity []AppFlowEvent) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "app_flow_event")
}

func (t *AppFlowEventRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "app_flow_event")
}

type AppFlowEvent struct {

	AppId      string  `json:"appId,omitempty"`

	Name      string  `json:"name,omitempty"`

	Aware      interface{}  `json:"aware,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	App      *App  `json:"app,omitempty"`

	Id      string  `json:"id,omitempty"`

	Flows      []Flow  `json:"flows,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}

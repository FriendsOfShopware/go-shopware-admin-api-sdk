package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type FlowRepository struct {
	*GenericRepository[Flow]
}

func NewFlowRepository(client *Client) *FlowRepository {
	return &FlowRepository{
		GenericRepository: NewGenericRepository[Flow](client),
	}
}

func (t *FlowRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[Flow], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "flow")
}

func (t *FlowRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[Flow], *http.Response, error) {
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

func (t *FlowRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "flow")
}

func (t *FlowRepository) Upsert(ctx ApiContext, entity []Flow) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "flow")
}

func (t *FlowRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "flow")
}

type Flow struct {

	Active      bool  `json:"active,omitempty"`

	AppFlowEvent      *AppFlowEvent  `json:"appFlowEvent,omitempty"`

	AppFlowEventId      string  `json:"appFlowEventId,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Description      string  `json:"description,omitempty"`

	EventName      string  `json:"eventName,omitempty"`

	Id      string  `json:"id,omitempty"`

	Invalid      bool  `json:"invalid,omitempty"`

	Name      string  `json:"name,omitempty"`

	Payload      interface{}  `json:"payload,omitempty"`

	Priority      float64  `json:"priority,omitempty"`

	Sequences      []FlowSequence  `json:"sequences,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}

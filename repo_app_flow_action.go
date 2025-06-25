package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type AppFlowActionRepository struct {
	*GenericRepository[AppFlowAction]
}

func NewAppFlowActionRepository(client *Client) *AppFlowActionRepository {
	return &AppFlowActionRepository{
		GenericRepository: NewGenericRepository[AppFlowAction](client),
	}
}

func (t *AppFlowActionRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[AppFlowAction], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "app-flow-action")
}

func (t *AppFlowActionRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[AppFlowAction], *http.Response, error) {
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

func (t *AppFlowActionRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "app-flow-action")
}

func (t *AppFlowActionRepository) Upsert(ctx ApiContext, entity []AppFlowAction) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "app_flow_action")
}

func (t *AppFlowActionRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "app_flow_action")
}

type AppFlowAction struct {

	App      *App  `json:"app,omitempty"`

	AppId      string  `json:"appId,omitempty"`

	Badge      string  `json:"badge,omitempty"`

	Config      interface{}  `json:"config,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Delayable      bool  `json:"delayable,omitempty"`

	Description      string  `json:"description,omitempty"`

	FlowSequences      []FlowSequence  `json:"flowSequences,omitempty"`

	Headers      interface{}  `json:"headers,omitempty"`

	Headline      string  `json:"headline,omitempty"`

	Icon      string  `json:"icon,omitempty"`

	IconRaw      interface{}  `json:"iconRaw,omitempty"`

	Id      string  `json:"id,omitempty"`

	Label      string  `json:"label,omitempty"`

	Name      string  `json:"name,omitempty"`

	Parameters      interface{}  `json:"parameters,omitempty"`

	Requirements      interface{}  `json:"requirements,omitempty"`

	SwIcon      string  `json:"swIcon,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	Translations      []AppFlowActionTranslation  `json:"translations,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Url      string  `json:"url,omitempty"`

}

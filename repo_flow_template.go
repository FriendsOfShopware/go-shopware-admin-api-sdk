package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type FlowTemplateRepository struct {
	*GenericRepository[FlowTemplate]
}

func NewFlowTemplateRepository(client *Client) *FlowTemplateRepository {
	return &FlowTemplateRepository{
		GenericRepository: NewGenericRepository[FlowTemplate](client),
	}
}

func (t *FlowTemplateRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[FlowTemplate], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "flow-template")
}

func (t *FlowTemplateRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[FlowTemplate], *http.Response, error) {
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

func (t *FlowTemplateRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "flow-template")
}

func (t *FlowTemplateRepository) Upsert(ctx ApiContext, entity []FlowTemplate) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "flow_template")
}

func (t *FlowTemplateRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "flow_template")
}

type FlowTemplate struct {

	Config      interface{}  `json:"config,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Id      string  `json:"id,omitempty"`

	Name      string  `json:"name,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}

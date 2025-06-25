package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type FlowSequenceRepository struct {
	*GenericRepository[FlowSequence]
}

func NewFlowSequenceRepository(client *Client) *FlowSequenceRepository {
	return &FlowSequenceRepository{
		GenericRepository: NewGenericRepository[FlowSequence](client),
	}
}

func (t *FlowSequenceRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[FlowSequence], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "flow-sequence")
}

func (t *FlowSequenceRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[FlowSequence], *http.Response, error) {
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

func (t *FlowSequenceRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "flow-sequence")
}

func (t *FlowSequenceRepository) Upsert(ctx ApiContext, entity []FlowSequence) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "flow_sequence")
}

func (t *FlowSequenceRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "flow_sequence")
}

type FlowSequence struct {

	Rule      *Rule  `json:"rule,omitempty"`

	Children      []FlowSequence  `json:"children,omitempty"`

	ParentId      string  `json:"parentId,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	FlowId      string  `json:"flowId,omitempty"`

	DisplayGroup      float64  `json:"displayGroup,omitempty"`

	TrueCase      bool  `json:"trueCase,omitempty"`

	Flow      *Flow  `json:"flow,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Id      string  `json:"id,omitempty"`

	Position      float64  `json:"position,omitempty"`

	AppFlowActionId      string  `json:"appFlowActionId,omitempty"`

	AppFlowAction      *AppFlowAction  `json:"appFlowAction,omitempty"`

	RuleId      string  `json:"ruleId,omitempty"`

	ActionName      string  `json:"actionName,omitempty"`

	Config      interface{}  `json:"config,omitempty"`

	Parent      *FlowSequence  `json:"parent,omitempty"`

}

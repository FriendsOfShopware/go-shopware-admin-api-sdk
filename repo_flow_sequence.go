package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type FlowSequenceRepository ClientService

func (t FlowSequenceRepository) Search(ctx ApiContext, criteria Criteria) (*FlowSequenceCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/flow-sequence", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(FlowSequenceCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t FlowSequenceRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/flow-sequence", criteria)

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

func (t FlowSequenceRepository) Upsert(ctx ApiContext, entity []FlowSequence) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"flow_sequence": {
		Entity:  "flow_sequence",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t FlowSequenceRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"flow_sequence": {
		Entity:  "flow_sequence",
		Action:  "delete",
		Payload: payload,
	}})
}

type FlowSequence struct {
	Config interface{} `json:"config,omitempty"`

	Position float64 `json:"position,omitempty"`

	Flow *Flow `json:"flow,omitempty"`

	Rule *Rule `json:"rule,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Id string `json:"id,omitempty"`

	FlowId string `json:"flowId,omitempty"`

	TrueCase bool `json:"trueCase,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	RuleId string `json:"ruleId,omitempty"`

	ActionName string `json:"actionName,omitempty"`

	DisplayGroup float64 `json:"displayGroup,omitempty"`

	Parent *FlowSequence `json:"parent,omitempty"`

	Children []FlowSequence `json:"children,omitempty"`

	ParentId string `json:"parentId,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`
}

type FlowSequenceCollection struct {
	EntityCollection

	Data []FlowSequence `json:"data"`
}

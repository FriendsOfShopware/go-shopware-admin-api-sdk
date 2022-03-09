package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type RuleConditionRepository ClientService

func (t RuleConditionRepository) Search(ctx ApiContext, criteria Criteria) (*RuleConditionCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/rule-condition", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(RuleConditionCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t RuleConditionRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/rule-condition", criteria)

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

func (t RuleConditionRepository) Upsert(ctx ApiContext, entity []RuleCondition) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"rule_condition": {
		Entity:  "rule_condition",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t RuleConditionRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"rule_condition": {
		Entity:  "rule_condition",
		Action:  "delete",
		Payload: payload,
	}})
}

type RuleCondition struct {
	Value interface{} `json:"value,omitempty"`

	Position float64 `json:"position,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Id string `json:"id,omitempty"`

	RuleId string `json:"ruleId,omitempty"`

	ParentId string `json:"parentId,omitempty"`

	Children []RuleCondition `json:"children,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Type string `json:"type,omitempty"`

	Rule *Rule `json:"rule,omitempty"`

	Parent *RuleCondition `json:"parent,omitempty"`
}

type RuleConditionCollection struct {
	EntityCollection

	Data []RuleCondition `json:"data"`
}

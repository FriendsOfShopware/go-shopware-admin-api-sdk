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

func (t RuleConditionRepository) SearchAll(ctx ApiContext, criteria Criteria) (*RuleConditionCollection, *http.Response, error) {
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

	Rule *Rule `json:"rule,omitempty"`

	Parent *RuleCondition `json:"parent,omitempty"`

	Position float64 `json:"position,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Id string `json:"id,omitempty"`

	Type string `json:"type,omitempty"`

	ScriptId string `json:"scriptId,omitempty"`

	ParentId string `json:"parentId,omitempty"`

	RuleId string `json:"ruleId,omitempty"`

	Children []RuleCondition `json:"children,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	AppScriptCondition *AppScriptCondition `json:"appScriptCondition,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type RuleConditionCollection struct {
	EntityCollection

	Data []RuleCondition `json:"data"`
}

package go_shopware_admin_sdk

import (
	"net/http"
)

type EventActionRuleRepository ClientService

func (t EventActionRuleRepository) Search(ctx ApiContext, criteria Criteria) (*EventActionRuleCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/event-action-rule", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(EventActionRuleCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t EventActionRuleRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/event-action-rule", criteria)

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

func (t EventActionRuleRepository) Upsert(ctx ApiContext, entity []EventActionRule) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"event_action_rule": {
		Entity:  "event_action_rule",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t EventActionRuleRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"event_action_rule": {
		Entity:  "event_action_rule",
		Action:  "delete",
		Payload: payload,
	}})
}

type EventActionRule struct {
	Rule *Rule `json:"rule,omitempty"`

	EventActionId string `json:"eventActionId,omitempty"`

	RuleId string `json:"ruleId,omitempty"`

	EventAction *EventAction `json:"eventAction,omitempty"`
}

type EventActionRuleCollection struct {
	EntityCollection

	Data []EventActionRule `json:"data"`
}

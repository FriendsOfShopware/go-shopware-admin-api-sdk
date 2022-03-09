package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type StateMachineTransitionRepository ClientService

func (t StateMachineTransitionRepository) Search(ctx ApiContext, criteria Criteria) (*StateMachineTransitionCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/state-machine-transition", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(StateMachineTransitionCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t StateMachineTransitionRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/state-machine-transition", criteria)

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

func (t StateMachineTransitionRepository) Upsert(ctx ApiContext, entity []StateMachineTransition) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"state_machine_transition": {
		Entity:  "state_machine_transition",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t StateMachineTransitionRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"state_machine_transition": {
		Entity:  "state_machine_transition",
		Action:  "delete",
		Payload: payload,
	}})
}

type StateMachineTransition struct {
	ActionName string `json:"actionName,omitempty"`

	StateMachine *StateMachine `json:"stateMachine,omitempty"`

	FromStateId string `json:"fromStateId,omitempty"`

	ToStateMachineState *StateMachineState `json:"toStateMachineState,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Id string `json:"id,omitempty"`

	StateMachineId string `json:"stateMachineId,omitempty"`

	FromStateMachineState *StateMachineState `json:"fromStateMachineState,omitempty"`

	ToStateId string `json:"toStateId,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type StateMachineTransitionCollection struct {
	EntityCollection

	Data []StateMachineTransition `json:"data"`
}

package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type StateMachineHistoryRepository ClientService

func (t StateMachineHistoryRepository) Search(ctx ApiContext, criteria Criteria) (*StateMachineHistoryCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/state-machine-history", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(StateMachineHistoryCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t StateMachineHistoryRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/state-machine-history", criteria)

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

func (t StateMachineHistoryRepository) Upsert(ctx ApiContext, entity []StateMachineHistory) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"state_machine_history": {
		Entity:  "state_machine_history",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t StateMachineHistoryRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"state_machine_history": {
		Entity:  "state_machine_history",
		Action:  "delete",
		Payload: payload,
	}})
}

type StateMachineHistory struct {
	FromStateMachineState *StateMachineState `json:"fromStateMachineState,omitempty"`

	UserId string `json:"userId,omitempty"`

	EntityId interface{} `json:"entityId,omitempty"`

	EntityName string `json:"entityName,omitempty"`

	FromStateId string `json:"fromStateId,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Id string `json:"id,omitempty"`

	StateMachine *StateMachine `json:"stateMachine,omitempty"`

	ToStateMachineState *StateMachineState `json:"toStateMachineState,omitempty"`

	TransitionActionName string `json:"transitionActionName,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	StateMachineId string `json:"stateMachineId,omitempty"`

	User *User `json:"user,omitempty"`

	ToStateId string `json:"toStateId,omitempty"`
}

type StateMachineHistoryCollection struct {
	EntityCollection

	Data []StateMachineHistory `json:"data"`
}

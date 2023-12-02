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

func (t StateMachineTransitionRepository) SearchAll(ctx ApiContext, criteria Criteria) (*StateMachineTransitionCollection, *http.Response, error) {
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

	Id      string  `json:"id,omitempty"`

	ActionName      string  `json:"actionName,omitempty"`

	StateMachine      *StateMachine  `json:"stateMachine,omitempty"`

	FromStateId      string  `json:"fromStateId,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	StateMachineId      string  `json:"stateMachineId,omitempty"`

	FromStateMachineState      *StateMachineState  `json:"fromStateMachineState,omitempty"`

	ToStateId      string  `json:"toStateId,omitempty"`

	ToStateMachineState      *StateMachineState  `json:"toStateMachineState,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

}

type StateMachineTransitionCollection struct {
	EntityCollection

	Data []StateMachineTransition `json:"data"`
}

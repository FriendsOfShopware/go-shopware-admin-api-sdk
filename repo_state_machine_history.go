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

func (t StateMachineHistoryRepository) SearchAll(ctx ApiContext, criteria Criteria) (*StateMachineHistoryCollection, *http.Response, error) {
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
	Id string `json:"id,omitempty"`

	FromStateId string `json:"fromStateId,omitempty"`

	ToStateId string `json:"toStateId,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	User *User `json:"user,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	ToStateMachineState *StateMachineState `json:"toStateMachineState,omitempty"`

	ReferencedVersionId string `json:"referencedVersionId,omitempty"`

	UserId string `json:"userId,omitempty"`

	ReferencedId string `json:"referencedId,omitempty"`

	StateMachineId string `json:"stateMachineId,omitempty"`

	StateMachine *StateMachine `json:"stateMachine,omitempty"`

	EntityName string `json:"entityName,omitempty"`

	FromStateMachineState *StateMachineState `json:"fromStateMachineState,omitempty"`

	TransitionActionName string `json:"transitionActionName,omitempty"`
}

type StateMachineHistoryCollection struct {
	EntityCollection

	Data []StateMachineHistory `json:"data"`
}

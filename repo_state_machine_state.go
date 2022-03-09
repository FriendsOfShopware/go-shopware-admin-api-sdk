package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type StateMachineStateRepository ClientService

func (t StateMachineStateRepository) Search(ctx ApiContext, criteria Criteria) (*StateMachineStateCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/state-machine-state", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(StateMachineStateCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t StateMachineStateRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/state-machine-state", criteria)

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

func (t StateMachineStateRepository) Upsert(ctx ApiContext, entity []StateMachineState) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"state_machine_state": {
		Entity:  "state_machine_state",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t StateMachineStateRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"state_machine_state": {
		Entity:  "state_machine_state",
		Action:  "delete",
		Payload: payload,
	}})
}

type StateMachineState struct {
	OrderTransactions []OrderTransaction `json:"orderTransactions,omitempty"`

	ToStateMachineHistoryEntries []StateMachineHistory `json:"toStateMachineHistoryEntries,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	OrderDeliveries []OrderDelivery `json:"orderDeliveries,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	StateMachine *StateMachine `json:"stateMachine,omitempty"`

	FromStateMachineTransitions []StateMachineTransition `json:"fromStateMachineTransitions,omitempty"`

	ToStateMachineTransitions []StateMachineTransition `json:"toStateMachineTransitions,omitempty"`

	Translations []StateMachineStateTranslation `json:"translations,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	TechnicalName string `json:"technicalName,omitempty"`

	StateMachineId string `json:"stateMachineId,omitempty"`

	Orders []Order `json:"orders,omitempty"`

	FromStateMachineHistoryEntries []StateMachineHistory `json:"fromStateMachineHistoryEntries,omitempty"`

	Translated interface{} `json:"translated,omitempty"`
}

type StateMachineStateCollection struct {
	EntityCollection

	Data []StateMachineState `json:"data"`
}

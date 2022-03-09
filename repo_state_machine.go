package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type StateMachineRepository ClientService

func (t StateMachineRepository) Search(ctx ApiContext, criteria Criteria) (*StateMachineCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/state-machine", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(StateMachineCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t StateMachineRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/state-machine", criteria)

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

func (t StateMachineRepository) Upsert(ctx ApiContext, entity []StateMachine) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"state_machine": {
		Entity:  "state_machine",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t StateMachineRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"state_machine": {
		Entity:  "state_machine",
		Action:  "delete",
		Payload: payload,
	}})
}

type StateMachine struct {
	Id string `json:"id,omitempty"`

	States []StateMachineState `json:"states,omitempty"`

	InitialStateId string `json:"initialStateId,omitempty"`

	HistoryEntries []StateMachineHistory `json:"historyEntries,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Translated interface{} `json:"translated,omitempty"`

	TechnicalName string `json:"technicalName,omitempty"`

	Name string `json:"name,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	Transitions []StateMachineTransition `json:"transitions,omitempty"`

	Translations []StateMachineTranslation `json:"translations,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`
}

type StateMachineCollection struct {
	EntityCollection

	Data []StateMachine `json:"data"`
}

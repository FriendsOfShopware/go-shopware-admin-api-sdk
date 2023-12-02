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

func (t StateMachineRepository) SearchAll(ctx ApiContext, criteria Criteria) (*StateMachineCollection, *http.Response, error) {
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
	InitialStateId string `json:"initialStateId,omitempty"`

	Translations []StateMachineTranslation `json:"translations,omitempty"`

	Name string `json:"name,omitempty"`

	TechnicalName string `json:"technicalName,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	States []StateMachineState `json:"states,omitempty"`

	Transitions []StateMachineTransition `json:"transitions,omitempty"`

	HistoryEntries []StateMachineHistory `json:"historyEntries,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Id string `json:"id,omitempty"`

	Translated interface{} `json:"translated,omitempty"`
}

type StateMachineCollection struct {
	EntityCollection

	Data []StateMachine `json:"data"`
}

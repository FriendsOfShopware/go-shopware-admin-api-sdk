package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type StateMachineStateTranslationRepository ClientService

func (t StateMachineStateTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*StateMachineStateTranslationCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/state-machine-state-translation", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(StateMachineStateTranslationCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t StateMachineStateTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/state-machine-state-translation", criteria)

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

func (t StateMachineStateTranslationRepository) Upsert(ctx ApiContext, entity []StateMachineStateTranslation) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"state_machine_state_translation": {
		Entity:  "state_machine_state_translation",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t StateMachineStateTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"state_machine_state_translation": {
		Entity:  "state_machine_state_translation",
		Action:  "delete",
		Payload: payload,
	}})
}

type StateMachineStateTranslation struct {
	StateMachineState *StateMachineState `json:"stateMachineState,omitempty"`

	Language *Language `json:"language,omitempty"`

	Name string `json:"name,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	StateMachineStateId string `json:"stateMachineStateId,omitempty"`

	LanguageId string `json:"languageId,omitempty"`
}

type StateMachineStateTranslationCollection struct {
	EntityCollection

	Data []StateMachineStateTranslation `json:"data"`
}

package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type StateMachineTranslationRepository ClientService

func (t StateMachineTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*StateMachineTranslationCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/state-machine-translation", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(StateMachineTranslationCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t StateMachineTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/state-machine-translation", criteria)

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

func (t StateMachineTranslationRepository) Upsert(ctx ApiContext, entity []StateMachineTranslation) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"state_machine_translation": {
		Entity:  "state_machine_translation",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t StateMachineTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"state_machine_translation": {
		Entity:  "state_machine_translation",
		Action:  "delete",
		Payload: payload,
	}})
}

type StateMachineTranslation struct {
	Language *Language `json:"language,omitempty"`

	Name string `json:"name,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	StateMachineId string `json:"stateMachineId,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	StateMachine *StateMachine `json:"stateMachine,omitempty"`
}

type StateMachineTranslationCollection struct {
	EntityCollection

	Data []StateMachineTranslation `json:"data"`
}

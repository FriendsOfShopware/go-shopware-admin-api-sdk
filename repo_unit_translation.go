package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type UnitTranslationRepository ClientService

func (t UnitTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*UnitTranslationCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/unit-translation", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(UnitTranslationCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t UnitTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/unit-translation", criteria)

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

func (t UnitTranslationRepository) Upsert(ctx ApiContext, entity []UnitTranslation) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"unit_translation": {
		Entity:  "unit_translation",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t UnitTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"unit_translation": {
		Entity:  "unit_translation",
		Action:  "delete",
		Payload: payload,
	}})
}

type UnitTranslation struct {
	ShortCode string `json:"shortCode,omitempty"`

	Name string `json:"name,omitempty"`

	Unit *Unit `json:"unit,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	UnitId string `json:"unitId,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	Language *Language `json:"language,omitempty"`
}

type UnitTranslationCollection struct {
	EntityCollection

	Data []UnitTranslation `json:"data"`
}

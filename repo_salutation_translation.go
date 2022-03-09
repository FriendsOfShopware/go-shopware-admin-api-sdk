package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type SalutationTranslationRepository ClientService

func (t SalutationTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*SalutationTranslationCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/salutation-translation", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(SalutationTranslationCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t SalutationTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/salutation-translation", criteria)

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

func (t SalutationTranslationRepository) Upsert(ctx ApiContext, entity []SalutationTranslation) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"salutation_translation": {
		Entity:  "salutation_translation",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t SalutationTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"salutation_translation": {
		Entity:  "salutation_translation",
		Action:  "delete",
		Payload: payload,
	}})
}

type SalutationTranslation struct {
	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	Language *Language `json:"language,omitempty"`

	DisplayName string `json:"displayName,omitempty"`

	LetterName string `json:"letterName,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	SalutationId string `json:"salutationId,omitempty"`

	Salutation *Salutation `json:"salutation,omitempty"`
}

type SalutationTranslationCollection struct {
	EntityCollection

	Data []SalutationTranslation `json:"data"`
}

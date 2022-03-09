package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type LocaleTranslationRepository ClientService

func (t LocaleTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*LocaleTranslationCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/locale-translation", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(LocaleTranslationCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t LocaleTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/locale-translation", criteria)

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

func (t LocaleTranslationRepository) Upsert(ctx ApiContext, entity []LocaleTranslation) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"locale_translation": {
		Entity:  "locale_translation",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t LocaleTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"locale_translation": {
		Entity:  "locale_translation",
		Action:  "delete",
		Payload: payload,
	}})
}

type LocaleTranslation struct {
	CustomFields interface{} `json:"customFields,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Language *Language `json:"language,omitempty"`

	LocaleId string `json:"localeId,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	Locale *Locale `json:"locale,omitempty"`

	Name string `json:"name,omitempty"`

	Territory string `json:"territory,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`
}

type LocaleTranslationCollection struct {
	EntityCollection

	Data []LocaleTranslation `json:"data"`
}

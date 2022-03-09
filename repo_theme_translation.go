package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type ThemeTranslationRepository ClientService

func (t ThemeTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*ThemeTranslationCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/theme-translation", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(ThemeTranslationCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ThemeTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/theme-translation", criteria)

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

func (t ThemeTranslationRepository) Upsert(ctx ApiContext, entity []ThemeTranslation) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"theme_translation": {
		Entity:  "theme_translation",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t ThemeTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"theme_translation": {
		Entity:  "theme_translation",
		Action:  "delete",
		Payload: payload,
	}})
}

type ThemeTranslation struct {
	Description string `json:"description,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Language *Language `json:"language,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	Theme *Theme `json:"theme,omitempty"`

	Labels interface{} `json:"labels,omitempty"`

	HelpTexts interface{} `json:"helpTexts,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	ThemeId string `json:"themeId,omitempty"`
}

type ThemeTranslationCollection struct {
	EntityCollection

	Data []ThemeTranslation `json:"data"`
}

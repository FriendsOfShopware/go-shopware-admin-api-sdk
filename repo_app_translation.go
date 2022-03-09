package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type AppTranslationRepository ClientService

func (t AppTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*AppTranslationCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/app-translation", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(AppTranslationCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t AppTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/app-translation", criteria)

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

func (t AppTranslationRepository) Upsert(ctx ApiContext, entity []AppTranslation) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"app_translation": {
		Entity:  "app_translation",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t AppTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"app_translation": {
		Entity:  "app_translation",
		Action:  "delete",
		Payload: payload,
	}})
}

type AppTranslation struct {
	Description string `json:"description,omitempty"`

	PrivacyPolicyExtensions string `json:"privacyPolicyExtensions,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	Language *Language `json:"language,omitempty"`

	Label string `json:"label,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	AppId string `json:"appId,omitempty"`

	App *App `json:"app,omitempty"`
}

type AppTranslationCollection struct {
	EntityCollection

	Data []AppTranslation `json:"data"`
}

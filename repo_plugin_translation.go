package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type PluginTranslationRepository ClientService

func (t PluginTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*PluginTranslationCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/plugin-translation", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(PluginTranslationCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t PluginTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/plugin-translation", criteria)

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

func (t PluginTranslationRepository) Upsert(ctx ApiContext, entity []PluginTranslation) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"plugin_translation": {
		Entity:  "plugin_translation",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t PluginTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"plugin_translation": {
		Entity:  "plugin_translation",
		Action:  "delete",
		Payload: payload,
	}})
}

type PluginTranslation struct {
	Description string `json:"description,omitempty"`

	SupportLink string `json:"supportLink,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	PluginId string `json:"pluginId,omitempty"`

	Plugin *Plugin `json:"plugin,omitempty"`

	Language *Language `json:"language,omitempty"`

	Label string `json:"label,omitempty"`

	ManufacturerLink string `json:"manufacturerLink,omitempty"`

	Changelog interface{} `json:"changelog,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	LanguageId string `json:"languageId,omitempty"`
}

type PluginTranslationCollection struct {
	EntityCollection

	Data []PluginTranslation `json:"data"`
}

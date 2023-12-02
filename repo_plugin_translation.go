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

func (t PluginTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*PluginTranslationCollection, *http.Response, error) {
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
	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Language *Language `json:"language,omitempty"`

	Label string `json:"label,omitempty"`

	Description string `json:"description,omitempty"`

	ManufacturerLink string `json:"manufacturerLink,omitempty"`

	SupportLink string `json:"supportLink,omitempty"`

	Changelog interface{} `json:"changelog,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	PluginId string `json:"pluginId,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	Plugin *Plugin `json:"plugin,omitempty"`
}

type PluginTranslationCollection struct {
	EntityCollection

	Data []PluginTranslation `json:"data"`
}

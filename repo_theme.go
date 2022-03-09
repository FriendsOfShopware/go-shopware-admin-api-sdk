package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type ThemeRepository ClientService

func (t ThemeRepository) Search(ctx ApiContext, criteria Criteria) (*ThemeCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/theme", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(ThemeCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ThemeRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/theme", criteria)

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

func (t ThemeRepository) Upsert(ctx ApiContext, entity []Theme) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"theme": {
		Entity:  "theme",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t ThemeRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"theme": {
		Entity:  "theme",
		Action:  "delete",
		Payload: payload,
	}})
}

type Theme struct {
	SalesChannels []SalesChannel `json:"salesChannels,omitempty"`

	ChildThemes []Theme `json:"childThemes,omitempty"`

	Media []Media `json:"media,omitempty"`

	PreviewMedia *Media `json:"previewMedia,omitempty"`

	Id string `json:"id,omitempty"`

	Labels interface{} `json:"labels,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	PreviewMediaId string `json:"previewMediaId,omitempty"`

	ParentThemeId string `json:"parentThemeId,omitempty"`

	Translations []ThemeTranslation `json:"translations,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Description string `json:"description,omitempty"`

	BaseConfig interface{} `json:"baseConfig,omitempty"`

	ConfigValues interface{} `json:"configValues,omitempty"`

	Active bool `json:"active,omitempty"`

	TechnicalName string `json:"technicalName,omitempty"`

	Name string `json:"name,omitempty"`

	Author string `json:"author,omitempty"`

	HelpTexts interface{} `json:"helpTexts,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Translated interface{} `json:"translated,omitempty"`
}

type ThemeCollection struct {
	EntityCollection

	Data []Theme `json:"data"`
}

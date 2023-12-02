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

func (t ThemeRepository) SearchAll(ctx ApiContext, criteria Criteria) (*ThemeCollection, *http.Response, error) {
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

	ConfigValues      interface{}  `json:"configValues,omitempty"`

	Translations      []ThemeTranslation  `json:"translations,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Labels      interface{}  `json:"labels,omitempty"`

	HelpTexts      interface{}  `json:"helpTexts,omitempty"`

	PreviewMediaId      string  `json:"previewMediaId,omitempty"`

	ParentThemeId      string  `json:"parentThemeId,omitempty"`

	ThemeJson      interface{}  `json:"themeJson,omitempty"`

	SalesChannels      []SalesChannel  `json:"salesChannels,omitempty"`

	DependentThemes      []Theme  `json:"dependentThemes,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Id      string  `json:"id,omitempty"`

	Author      string  `json:"author,omitempty"`

	BaseConfig      interface{}  `json:"baseConfig,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	Name      string  `json:"name,omitempty"`

	Description      string  `json:"description,omitempty"`

	Active      bool  `json:"active,omitempty"`

	Media      []Media  `json:"media,omitempty"`

	PreviewMedia      *Media  `json:"previewMedia,omitempty"`

	TechnicalName      string  `json:"technicalName,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

}

type ThemeCollection struct {
	EntityCollection

	Data []Theme `json:"data"`
}

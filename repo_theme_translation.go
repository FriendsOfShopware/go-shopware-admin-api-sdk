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

func (t ThemeTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*ThemeTranslationCollection, *http.Response, error) {
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

	Language      *Language  `json:"language,omitempty"`

	Labels      interface{}  `json:"labels,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	ThemeId      string  `json:"themeId,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	Theme      *Theme  `json:"theme,omitempty"`

	Description      string  `json:"description,omitempty"`

	HelpTexts      interface{}  `json:"helpTexts,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

}

type ThemeTranslationCollection struct {
	EntityCollection

	Data []ThemeTranslation `json:"data"`
}

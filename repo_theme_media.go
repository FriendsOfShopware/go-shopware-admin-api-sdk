package go_shopware_admin_sdk

import (
	"net/http"
)

type ThemeMediaRepository ClientService

func (t ThemeMediaRepository) Search(ctx ApiContext, criteria Criteria) (*ThemeMediaCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/theme-media", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(ThemeMediaCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ThemeMediaRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/theme-media", criteria)

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

func (t ThemeMediaRepository) Upsert(ctx ApiContext, entity []ThemeMedia) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"theme_media": {
		Entity:  "theme_media",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t ThemeMediaRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"theme_media": {
		Entity:  "theme_media",
		Action:  "delete",
		Payload: payload,
	}})
}

type ThemeMedia struct {
	ThemeId string `json:"themeId,omitempty"`

	MediaId string `json:"mediaId,omitempty"`

	Theme *Theme `json:"theme,omitempty"`

	Media *Media `json:"media,omitempty"`
}

type ThemeMediaCollection struct {
	EntityCollection

	Data []ThemeMedia `json:"data"`
}

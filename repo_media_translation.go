package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type MediaTranslationRepository ClientService

func (t MediaTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*MediaTranslationCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/media-translation", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(MediaTranslationCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t MediaTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/media-translation", criteria)

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

func (t MediaTranslationRepository) Upsert(ctx ApiContext, entity []MediaTranslation) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"media_translation": {
		Entity:  "media_translation",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t MediaTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"media_translation": {
		Entity:  "media_translation",
		Action:  "delete",
		Payload: payload,
	}})
}

type MediaTranslation struct {
	LanguageId string `json:"languageId,omitempty"`

	Language *Language `json:"language,omitempty"`

	Alt string `json:"alt,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	MediaId string `json:"mediaId,omitempty"`

	Media *Media `json:"media,omitempty"`

	Title string `json:"title,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type MediaTranslationCollection struct {
	EntityCollection

	Data []MediaTranslation `json:"data"`
}

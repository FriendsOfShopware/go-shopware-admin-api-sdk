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

func (t MediaTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*MediaTranslationCollection, *http.Response, error) {
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

	Media *Media `json:"media,omitempty"`

	Title string `json:"title,omitempty"`

	Alt string `json:"alt,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	MediaId string `json:"mediaId,omitempty"`

	Language *Language `json:"language,omitempty"`
}

type MediaTranslationCollection struct {
	EntityCollection

	Data []MediaTranslation `json:"data"`
}

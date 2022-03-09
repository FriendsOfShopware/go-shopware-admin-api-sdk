package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type CmsPageTranslationRepository ClientService

func (t CmsPageTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*CmsPageTranslationCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/cms-page-translation", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(CmsPageTranslationCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t CmsPageTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/cms-page-translation", criteria)

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

func (t CmsPageTranslationRepository) Upsert(ctx ApiContext, entity []CmsPageTranslation) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"cms_page_translation": {
		Entity:  "cms_page_translation",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t CmsPageTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"cms_page_translation": {
		Entity:  "cms_page_translation",
		Action:  "delete",
		Payload: payload,
	}})
}

type CmsPageTranslation struct {
	CmsPageId string `json:"cmsPageId,omitempty"`

	Language *Language `json:"language,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	CmsPageVersionId string `json:"cmsPageVersionId,omitempty"`

	Name string `json:"name,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	CmsPage *CmsPage `json:"cmsPage,omitempty"`
}

type CmsPageTranslationCollection struct {
	EntityCollection

	Data []CmsPageTranslation `json:"data"`
}

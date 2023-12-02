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

func (t CmsPageTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*CmsPageTranslationCollection, *http.Response, error) {
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

	Name      string  `json:"name,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CmsPageId      string  `json:"cmsPageId,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	CmsPage      *CmsPage  `json:"cmsPage,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	CmsPageVersionId      string  `json:"cmsPageVersionId,omitempty"`

}

type CmsPageTranslationCollection struct {
	EntityCollection

	Data []CmsPageTranslation `json:"data"`
}

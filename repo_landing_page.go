package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type LandingPageRepository ClientService

func (t LandingPageRepository) Search(ctx ApiContext, criteria Criteria) (*LandingPageCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/landing-page", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(LandingPageCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t LandingPageRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/landing-page", criteria)

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

func (t LandingPageRepository) Upsert(ctx ApiContext, entity []LandingPage) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"landing_page": {
		Entity:  "landing_page",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t LandingPageRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"landing_page": {
		Entity:  "landing_page",
		Action:  "delete",
		Payload: payload,
	}})
}

type LandingPage struct {
	SeoUrls []SeoUrl `json:"seoUrls,omitempty"`

	Id string `json:"id,omitempty"`

	VersionId string `json:"versionId,omitempty"`

	Active bool `json:"active,omitempty"`

	Name string `json:"name,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	SlotConfig interface{} `json:"slotConfig,omitempty"`

	Translations []LandingPageTranslation `json:"translations,omitempty"`

	MetaTitle string `json:"metaTitle,omitempty"`

	MetaDescription string `json:"metaDescription,omitempty"`

	Keywords string `json:"keywords,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Tags []Tag `json:"tags,omitempty"`

	SalesChannels []SalesChannel `json:"salesChannels,omitempty"`

	CmsPageVersionId string `json:"cmsPageVersionId,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Url string `json:"url,omitempty"`

	CmsPageId string `json:"cmsPageId,omitempty"`

	CmsPage *CmsPage `json:"cmsPage,omitempty"`

	Translated interface{} `json:"translated,omitempty"`
}

type LandingPageCollection struct {
	EntityCollection

	Data []LandingPage `json:"data"`
}

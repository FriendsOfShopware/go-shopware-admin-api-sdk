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

func (t LandingPageRepository) SearchAll(ctx ApiContext, criteria Criteria) (*LandingPageCollection, *http.Response, error) {
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
	Name string `json:"name,omitempty"`

	SalesChannels []SalesChannel `json:"salesChannels,omitempty"`

	CmsPageVersionId string `json:"cmsPageVersionId,omitempty"`

	SlotConfig interface{} `json:"slotConfig,omitempty"`

	Translations []LandingPageTranslation `json:"translations,omitempty"`

	Tags []Tag `json:"tags,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Translated interface{} `json:"translated,omitempty"`

	Id string `json:"id,omitempty"`

	VersionId string `json:"versionId,omitempty"`

	MetaTitle string `json:"metaTitle,omitempty"`

	MetaDescription string `json:"metaDescription,omitempty"`

	Url string `json:"url,omitempty"`

	CmsPageId string `json:"cmsPageId,omitempty"`

	Active bool `json:"active,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	Keywords string `json:"keywords,omitempty"`

	CmsPage *CmsPage `json:"cmsPage,omitempty"`

	SeoUrls []SeoUrl `json:"seoUrls,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type LandingPageCollection struct {
	EntityCollection

	Data []LandingPage `json:"data"`
}

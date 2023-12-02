package go_shopware_admin_sdk

import (
	"net/http"

	"time"
)

type CmsPageRepository ClientService

func (t CmsPageRepository) Search(ctx ApiContext, criteria Criteria) (*CmsPageCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/cms-page", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(CmsPageCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t CmsPageRepository) SearchAll(ctx ApiContext, criteria Criteria) (*CmsPageCollection, *http.Response, error) {
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

func (t CmsPageRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/cms-page", criteria)

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

func (t CmsPageRepository) Upsert(ctx ApiContext, entity []CmsPage) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"cms_page": {
		Entity:  "cms_page",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t CmsPageRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"cms_page": {
		Entity:  "cms_page",
		Action:  "delete",
		Payload: payload,
	}})
}

type CmsPage struct {
	Config interface{} `json:"config,omitempty"`

	Translations []CmsPageTranslation `json:"translations,omitempty"`

	Translated interface{} `json:"translated,omitempty"`

	Name string `json:"name,omitempty"`

	CssClass string `json:"cssClass,omitempty"`

	Type string `json:"type,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Id string `json:"id,omitempty"`

	VersionId string `json:"versionId,omitempty"`

	PreviewMedia *Media `json:"previewMedia,omitempty"`

	Categories []Category `json:"categories,omitempty"`

	LandingPages []LandingPage `json:"landingPages,omitempty"`

	HomeSalesChannels []SalesChannel `json:"homeSalesChannels,omitempty"`

	Entity string `json:"entity,omitempty"`

	PreviewMediaId string `json:"previewMediaId,omitempty"`

	Products []Product `json:"products,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Locked bool `json:"locked,omitempty"`

	Sections []CmsSection `json:"sections,omitempty"`
}

type CmsPageCollection struct {
	EntityCollection

	Data []CmsPage `json:"data"`
}

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
	Sections []CmsSection `json:"sections,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Products []Product `json:"products,omitempty"`

	Id string `json:"id,omitempty"`

	VersionId string `json:"versionId,omitempty"`

	Config interface{} `json:"config,omitempty"`

	PreviewMediaId string `json:"previewMediaId,omitempty"`

	Translations []CmsPageTranslation `json:"translations,omitempty"`

	LandingPages []LandingPage `json:"landingPages,omitempty"`

	Entity string `json:"entity,omitempty"`

	HomeSalesChannels []SalesChannel `json:"homeSalesChannels,omitempty"`

	Translated interface{} `json:"translated,omitempty"`

	Name string `json:"name,omitempty"`

	Type string `json:"type,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	Locked bool `json:"locked,omitempty"`

	PreviewMedia *Media `json:"previewMedia,omitempty"`

	Categories []Category `json:"categories,omitempty"`
}

type CmsPageCollection struct {
	EntityCollection

	Data []CmsPage `json:"data"`
}

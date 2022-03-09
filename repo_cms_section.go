package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type CmsSectionRepository ClientService

func (t CmsSectionRepository) Search(ctx ApiContext, criteria Criteria) (*CmsSectionCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/cms-section", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(CmsSectionCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t CmsSectionRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/cms-section", criteria)

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

func (t CmsSectionRepository) Upsert(ctx ApiContext, entity []CmsSection) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"cms_section": {
		Entity:  "cms_section",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t CmsSectionRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"cms_section": {
		Entity:  "cms_section",
		Action:  "delete",
		Payload: payload,
	}})
}

type CmsSection struct {
	Type string `json:"type,omitempty"`

	MobileBehavior string `json:"mobileBehavior,omitempty"`

	PageId string `json:"pageId,omitempty"`

	Page *CmsPage `json:"page,omitempty"`

	CssClass string `json:"cssClass,omitempty"`

	BackgroundMedia *Media `json:"backgroundMedia,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	VersionId string `json:"versionId,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Id string `json:"id,omitempty"`

	Position float64 `json:"position,omitempty"`

	SizingMode string `json:"sizingMode,omitempty"`

	BackgroundColor string `json:"backgroundColor,omitempty"`

	BackgroundMediaMode string `json:"backgroundMediaMode,omitempty"`

	CmsPageVersionId string `json:"cmsPageVersionId,omitempty"`

	Locked bool `json:"locked,omitempty"`

	Name string `json:"name,omitempty"`

	BackgroundMediaId string `json:"backgroundMediaId,omitempty"`

	Blocks []CmsBlock `json:"blocks,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type CmsSectionCollection struct {
	EntityCollection

	Data []CmsSection `json:"data"`
}

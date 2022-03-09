package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type CmsBlockRepository ClientService

func (t CmsBlockRepository) Search(ctx ApiContext, criteria Criteria) (*CmsBlockCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/cms-block", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(CmsBlockCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t CmsBlockRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/cms-block", criteria)

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

func (t CmsBlockRepository) Upsert(ctx ApiContext, entity []CmsBlock) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"cms_block": {
		Entity:  "cms_block",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t CmsBlockRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"cms_block": {
		Entity:  "cms_block",
		Action:  "delete",
		Payload: payload,
	}})
}

type CmsBlock struct {
	Section *CmsSection `json:"section,omitempty"`

	CmsSectionVersionId string `json:"cmsSectionVersionId,omitempty"`

	SectionPosition string `json:"sectionPosition,omitempty"`

	Type string `json:"type,omitempty"`

	Name string `json:"name,omitempty"`

	MarginBottom string `json:"marginBottom,omitempty"`

	MarginRight string `json:"marginRight,omitempty"`

	CssClass string `json:"cssClass,omitempty"`

	SectionId string `json:"sectionId,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	Id string `json:"id,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	MarginTop string `json:"marginTop,omitempty"`

	MarginLeft string `json:"marginLeft,omitempty"`

	BackgroundColor string `json:"backgroundColor,omitempty"`

	BackgroundMediaId string `json:"backgroundMediaId,omitempty"`

	VersionId string `json:"versionId,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Locked bool `json:"locked,omitempty"`

	BackgroundMediaMode string `json:"backgroundMediaMode,omitempty"`

	BackgroundMedia *Media `json:"backgroundMedia,omitempty"`

	Slots []CmsSlot `json:"slots,omitempty"`

	Position float64 `json:"position,omitempty"`
}

type CmsBlockCollection struct {
	EntityCollection

	Data []CmsBlock `json:"data"`
}

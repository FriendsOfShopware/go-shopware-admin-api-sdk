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

func (t CmsBlockRepository) SearchAll(ctx ApiContext, criteria Criteria) (*CmsBlockCollection, *http.Response, error) {
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
	SectionPosition string `json:"sectionPosition,omitempty"`

	MarginBottom string `json:"marginBottom,omitempty"`

	CssClass string `json:"cssClass,omitempty"`

	BackgroundMedia *Media `json:"backgroundMedia,omitempty"`

	SectionId string `json:"sectionId,omitempty"`

	Section *CmsSection `json:"section,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	Id string `json:"id,omitempty"`

	Position float64 `json:"position,omitempty"`

	MarginTop string `json:"marginTop,omitempty"`

	BackgroundMediaMode string `json:"backgroundMediaMode,omitempty"`

	MarginLeft string `json:"marginLeft,omitempty"`

	BackgroundColor string `json:"backgroundColor,omitempty"`

	BackgroundMediaId string `json:"backgroundMediaId,omitempty"`

	Slots []CmsSlot `json:"slots,omitempty"`

	VersionId string `json:"versionId,omitempty"`

	CmsSectionVersionId string `json:"cmsSectionVersionId,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Type string `json:"type,omitempty"`

	Locked bool `json:"locked,omitempty"`

	Name string `json:"name,omitempty"`

	MarginRight string `json:"marginRight,omitempty"`
}

type CmsBlockCollection struct {
	EntityCollection

	Data []CmsBlock `json:"data"`
}

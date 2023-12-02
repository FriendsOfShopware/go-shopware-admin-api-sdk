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

	Visibility      interface{}  `json:"visibility,omitempty"`

	BackgroundMedia      *Media  `json:"backgroundMedia,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Locked      bool  `json:"locked,omitempty"`

	Name      string  `json:"name,omitempty"`

	MarginTop      string  `json:"marginTop,omitempty"`

	MarginLeft      string  `json:"marginLeft,omitempty"`

	MarginRight      string  `json:"marginRight,omitempty"`

	Id      string  `json:"id,omitempty"`

	BackgroundMediaId      string  `json:"backgroundMediaId,omitempty"`

	Section      *CmsSection  `json:"section,omitempty"`

	VersionId      string  `json:"versionId,omitempty"`

	CmsSectionVersionId      string  `json:"cmsSectionVersionId,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Type      string  `json:"type,omitempty"`

	MarginBottom      string  `json:"marginBottom,omitempty"`

	BackgroundColor      string  `json:"backgroundColor,omitempty"`

	SectionId      string  `json:"sectionId,omitempty"`

	Slots      []CmsSlot  `json:"slots,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Position      float64  `json:"position,omitempty"`

	SectionPosition      string  `json:"sectionPosition,omitempty"`

	BackgroundMediaMode      string  `json:"backgroundMediaMode,omitempty"`

	CssClass      string  `json:"cssClass,omitempty"`

}

type CmsBlockCollection struct {
	EntityCollection

	Data []CmsBlock `json:"data"`
}

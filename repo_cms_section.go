package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type CmsSectionRepository struct {
	*GenericRepository[CmsSection]
}

func NewCmsSectionRepository(client *Client) *CmsSectionRepository {
	return &CmsSectionRepository{
		GenericRepository: NewGenericRepository[CmsSection](client),
	}
}

func (t *CmsSectionRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[CmsSection], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "cms-section")
}

func (t *CmsSectionRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[CmsSection], *http.Response, error) {
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

func (t *CmsSectionRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "cms-section")
}

func (t *CmsSectionRepository) Upsert(ctx ApiContext, entity []CmsSection) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "cms_section")
}

func (t *CmsSectionRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "cms_section")
}

type CmsSection struct {

	BackgroundMedia      *Media  `json:"backgroundMedia,omitempty"`

	Id      string  `json:"id,omitempty"`

	Position      float64  `json:"position,omitempty"`

	Locked      bool  `json:"locked,omitempty"`

	MobileBehavior      string  `json:"mobileBehavior,omitempty"`

	PageId      string  `json:"pageId,omitempty"`

	Blocks      []CmsBlock  `json:"blocks,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Name      string  `json:"name,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	CmsPageVersionId      string  `json:"cmsPageVersionId,omitempty"`

	BackgroundMediaMode      string  `json:"backgroundMediaMode,omitempty"`

	VersionId      string  `json:"versionId,omitempty"`

	Type      string  `json:"type,omitempty"`

	SizingMode      string  `json:"sizingMode,omitempty"`

	BackgroundColor      string  `json:"backgroundColor,omitempty"`

	BackgroundMediaId      string  `json:"backgroundMediaId,omitempty"`

	CssClass      string  `json:"cssClass,omitempty"`

	Visibility      interface{}  `json:"visibility,omitempty"`

	Page      *CmsPage  `json:"page,omitempty"`

}

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

	BackgroundColor      string  `json:"backgroundColor,omitempty"`

	BackgroundMedia      *Media  `json:"backgroundMedia,omitempty"`

	BackgroundMediaId      string  `json:"backgroundMediaId,omitempty"`

	BackgroundMediaMode      string  `json:"backgroundMediaMode,omitempty"`

	Blocks      []CmsBlock  `json:"blocks,omitempty"`

	CmsPageVersionId      string  `json:"cmsPageVersionId,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CssClass      string  `json:"cssClass,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Id      string  `json:"id,omitempty"`

	Locked      bool  `json:"locked,omitempty"`

	MobileBehavior      string  `json:"mobileBehavior,omitempty"`

	Name      string  `json:"name,omitempty"`

	Page      *CmsPage  `json:"page,omitempty"`

	PageId      string  `json:"pageId,omitempty"`

	Position      float64  `json:"position,omitempty"`

	SizingMode      string  `json:"sizingMode,omitempty"`

	Type      string  `json:"type,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	VersionId      string  `json:"versionId,omitempty"`

	Visibility      interface{}  `json:"visibility,omitempty"`

}

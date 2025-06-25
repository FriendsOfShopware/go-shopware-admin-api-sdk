package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type CmsBlockRepository struct {
	*GenericRepository[CmsBlock]
}

func NewCmsBlockRepository(client *Client) *CmsBlockRepository {
	return &CmsBlockRepository{
		GenericRepository: NewGenericRepository[CmsBlock](client),
	}
}

func (t *CmsBlockRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[CmsBlock], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "cms-block")
}

func (t *CmsBlockRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[CmsBlock], *http.Response, error) {
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

func (t *CmsBlockRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "cms-block")
}

func (t *CmsBlockRepository) Upsert(ctx ApiContext, entity []CmsBlock) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "cms_block")
}

func (t *CmsBlockRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "cms_block")
}

type CmsBlock struct {

	BackgroundColor      string  `json:"backgroundColor,omitempty"`

	BackgroundMedia      *Media  `json:"backgroundMedia,omitempty"`

	BackgroundMediaId      string  `json:"backgroundMediaId,omitempty"`

	BackgroundMediaMode      string  `json:"backgroundMediaMode,omitempty"`

	CmsSectionVersionId      string  `json:"cmsSectionVersionId,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CssClass      string  `json:"cssClass,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Id      string  `json:"id,omitempty"`

	Locked      bool  `json:"locked,omitempty"`

	MarginBottom      string  `json:"marginBottom,omitempty"`

	MarginLeft      string  `json:"marginLeft,omitempty"`

	MarginRight      string  `json:"marginRight,omitempty"`

	MarginTop      string  `json:"marginTop,omitempty"`

	Name      string  `json:"name,omitempty"`

	Position      float64  `json:"position,omitempty"`

	Section      *CmsSection  `json:"section,omitempty"`

	SectionId      string  `json:"sectionId,omitempty"`

	SectionPosition      string  `json:"sectionPosition,omitempty"`

	Slots      []CmsSlot  `json:"slots,omitempty"`

	Type      string  `json:"type,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	VersionId      string  `json:"versionId,omitempty"`

	Visibility      interface{}  `json:"visibility,omitempty"`

}

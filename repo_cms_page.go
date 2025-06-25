package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type CmsPageRepository struct {
	*GenericRepository[CmsPage]
}

func NewCmsPageRepository(client *Client) *CmsPageRepository {
	return &CmsPageRepository{
		GenericRepository: NewGenericRepository[CmsPage](client),
	}
}

func (t *CmsPageRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[CmsPage], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "cms-page")
}

func (t *CmsPageRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[CmsPage], *http.Response, error) {
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

func (t *CmsPageRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "cms-page")
}

func (t *CmsPageRepository) Upsert(ctx ApiContext, entity []CmsPage) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "cms_page")
}

func (t *CmsPageRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "cms_page")
}

type CmsPage struct {

	Categories      []Category  `json:"categories,omitempty"`

	Config      interface{}  `json:"config,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CssClass      string  `json:"cssClass,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Entity      string  `json:"entity,omitempty"`

	HomeSalesChannels      []SalesChannel  `json:"homeSalesChannels,omitempty"`

	Id      string  `json:"id,omitempty"`

	LandingPages      []LandingPage  `json:"landingPages,omitempty"`

	Locked      bool  `json:"locked,omitempty"`

	Name      string  `json:"name,omitempty"`

	PreviewMedia      *Media  `json:"previewMedia,omitempty"`

	PreviewMediaId      string  `json:"previewMediaId,omitempty"`

	Products      []Product  `json:"products,omitempty"`

	Sections      []CmsSection  `json:"sections,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	Translations      []CmsPageTranslation  `json:"translations,omitempty"`

	Type      string  `json:"type,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	VersionId      string  `json:"versionId,omitempty"`

}

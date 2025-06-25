package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type LandingPageRepository struct {
	*GenericRepository[LandingPage]
}

func NewLandingPageRepository(client *Client) *LandingPageRepository {
	return &LandingPageRepository{
		GenericRepository: NewGenericRepository[LandingPage](client),
	}
}

func (t *LandingPageRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[LandingPage], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "landing-page")
}

func (t *LandingPageRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[LandingPage], *http.Response, error) {
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

func (t *LandingPageRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "landing-page")
}

func (t *LandingPageRepository) Upsert(ctx ApiContext, entity []LandingPage) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "landing_page")
}

func (t *LandingPageRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "landing_page")
}

type LandingPage struct {

	Active      bool  `json:"active,omitempty"`

	CmsPage      *CmsPage  `json:"cmsPage,omitempty"`

	CmsPageId      string  `json:"cmsPageId,omitempty"`

	CmsPageVersionId      string  `json:"cmsPageVersionId,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Id      string  `json:"id,omitempty"`

	Keywords      string  `json:"keywords,omitempty"`

	MetaDescription      string  `json:"metaDescription,omitempty"`

	MetaTitle      string  `json:"metaTitle,omitempty"`

	Name      string  `json:"name,omitempty"`

	SalesChannels      []SalesChannel  `json:"salesChannels,omitempty"`

	SeoUrls      []SeoUrl  `json:"seoUrls,omitempty"`

	SlotConfig      interface{}  `json:"slotConfig,omitempty"`

	Tags      []Tag  `json:"tags,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	Translations      []LandingPageTranslation  `json:"translations,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Url      string  `json:"url,omitempty"`

	VersionId      string  `json:"versionId,omitempty"`

}

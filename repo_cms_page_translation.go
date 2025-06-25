package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type CmsPageTranslationRepository struct {
	*GenericRepository[CmsPageTranslation]
}

func NewCmsPageTranslationRepository(client *Client) *CmsPageTranslationRepository {
	return &CmsPageTranslationRepository{
		GenericRepository: NewGenericRepository[CmsPageTranslation](client),
	}
}

func (t *CmsPageTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[CmsPageTranslation], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "cms-page-translation")
}

func (t *CmsPageTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[CmsPageTranslation], *http.Response, error) {
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

func (t *CmsPageTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "cms-page-translation")
}

func (t *CmsPageTranslationRepository) Upsert(ctx ApiContext, entity []CmsPageTranslation) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "cms_page_translation")
}

func (t *CmsPageTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "cms_page_translation")
}

type CmsPageTranslation struct {

	CmsPage      *CmsPage  `json:"cmsPage,omitempty"`

	CmsPageId      string  `json:"cmsPageId,omitempty"`

	CmsPageVersionId      string  `json:"cmsPageVersionId,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	Name      string  `json:"name,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}

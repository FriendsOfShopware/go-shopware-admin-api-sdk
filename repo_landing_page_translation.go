package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type LandingPageTranslationRepository struct {
	*GenericRepository[LandingPageTranslation]
}

func NewLandingPageTranslationRepository(client *Client) *LandingPageTranslationRepository {
	return &LandingPageTranslationRepository{
		GenericRepository: NewGenericRepository[LandingPageTranslation](client),
	}
}

func (t *LandingPageTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[LandingPageTranslation], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "landing-page-translation")
}

func (t *LandingPageTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[LandingPageTranslation], *http.Response, error) {
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

func (t *LandingPageTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "landing-page-translation")
}

func (t *LandingPageTranslationRepository) Upsert(ctx ApiContext, entity []LandingPageTranslation) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "landing_page_translation")
}

func (t *LandingPageTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "landing_page_translation")
}

type LandingPageTranslation struct {

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Keywords      string  `json:"keywords,omitempty"`

	LandingPage      *LandingPage  `json:"landingPage,omitempty"`

	LandingPageId      string  `json:"landingPageId,omitempty"`

	LandingPageVersionId      string  `json:"landingPageVersionId,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	MetaDescription      string  `json:"metaDescription,omitempty"`

	MetaTitle      string  `json:"metaTitle,omitempty"`

	Name      string  `json:"name,omitempty"`

	SlotConfig      interface{}  `json:"slotConfig,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Url      string  `json:"url,omitempty"`

}

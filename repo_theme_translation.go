package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type ThemeTranslationRepository struct {
	*GenericRepository[ThemeTranslation]
}

func NewThemeTranslationRepository(client *Client) *ThemeTranslationRepository {
	return &ThemeTranslationRepository{
		GenericRepository: NewGenericRepository[ThemeTranslation](client),
	}
}

func (t *ThemeTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[ThemeTranslation], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "theme-translation")
}

func (t *ThemeTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[ThemeTranslation], *http.Response, error) {
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

func (t *ThemeTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "theme-translation")
}

func (t *ThemeTranslationRepository) Upsert(ctx ApiContext, entity []ThemeTranslation) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "theme_translation")
}

func (t *ThemeTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "theme_translation")
}

type ThemeTranslation struct {

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Description      string  `json:"description,omitempty"`

	HelpTexts      interface{}  `json:"helpTexts,omitempty"`

	Labels      interface{}  `json:"labels,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	Theme      *Theme  `json:"theme,omitempty"`

	ThemeId      string  `json:"themeId,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}

package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type ThemeRepository struct {
	*GenericRepository[Theme]
}

func NewThemeRepository(client *Client) *ThemeRepository {
	return &ThemeRepository{
		GenericRepository: NewGenericRepository[Theme](client),
	}
}

func (t *ThemeRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[Theme], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "theme")
}

func (t *ThemeRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[Theme], *http.Response, error) {
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

func (t *ThemeRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "theme")
}

func (t *ThemeRepository) Upsert(ctx ApiContext, entity []Theme) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "theme")
}

func (t *ThemeRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "theme")
}

type Theme struct {

	Active      bool  `json:"active,omitempty"`

	Author      string  `json:"author,omitempty"`

	BaseConfig      interface{}  `json:"baseConfig,omitempty"`

	ConfigValues      interface{}  `json:"configValues,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	DependentThemes      []Theme  `json:"dependentThemes,omitempty"`

	Description      string  `json:"description,omitempty"`

	HelpTexts      interface{}  `json:"helpTexts,omitempty"`

	Id      string  `json:"id,omitempty"`

	Labels      interface{}  `json:"labels,omitempty"`

	Media      []Media  `json:"media,omitempty"`

	Name      string  `json:"name,omitempty"`

	ParentThemeId      string  `json:"parentThemeId,omitempty"`

	PreviewMedia      *Media  `json:"previewMedia,omitempty"`

	PreviewMediaId      string  `json:"previewMediaId,omitempty"`

	SalesChannels      []SalesChannel  `json:"salesChannels,omitempty"`

	TechnicalName      string  `json:"technicalName,omitempty"`

	ThemeJson      interface{}  `json:"themeJson,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	Translations      []ThemeTranslation  `json:"translations,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}

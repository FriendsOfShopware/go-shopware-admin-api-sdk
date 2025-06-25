package go_shopware_admin_sdk

import (
	"net/http"

)

type ThemeMediaRepository struct {
	*GenericRepository[ThemeMedia]
}

func NewThemeMediaRepository(client *Client) *ThemeMediaRepository {
	return &ThemeMediaRepository{
		GenericRepository: NewGenericRepository[ThemeMedia](client),
	}
}

func (t *ThemeMediaRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[ThemeMedia], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "theme-media")
}

func (t *ThemeMediaRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[ThemeMedia], *http.Response, error) {
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

func (t *ThemeMediaRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "theme-media")
}

func (t *ThemeMediaRepository) Upsert(ctx ApiContext, entity []ThemeMedia) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "theme_media")
}

func (t *ThemeMediaRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "theme_media")
}

type ThemeMedia struct {

	Media      *Media  `json:"media,omitempty"`

	MediaId      string  `json:"mediaId,omitempty"`

	Theme      *Theme  `json:"theme,omitempty"`

	ThemeId      string  `json:"themeId,omitempty"`

}

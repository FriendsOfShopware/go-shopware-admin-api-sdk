package go_shopware_admin_sdk

import (
	"net/http"

)

type ThemeChildRepository struct {
	*GenericRepository[ThemeChild]
}

func NewThemeChildRepository(client *Client) *ThemeChildRepository {
	return &ThemeChildRepository{
		GenericRepository: NewGenericRepository[ThemeChild](client),
	}
}

func (t *ThemeChildRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[ThemeChild], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "theme-child")
}

func (t *ThemeChildRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[ThemeChild], *http.Response, error) {
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

func (t *ThemeChildRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "theme-child")
}

func (t *ThemeChildRepository) Upsert(ctx ApiContext, entity []ThemeChild) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "theme_child")
}

func (t *ThemeChildRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "theme_child")
}

type ThemeChild struct {

	ChildTheme      *Theme  `json:"childTheme,omitempty"`

	ParentId      string  `json:"parentId,omitempty"`

	ChildId      string  `json:"childId,omitempty"`

	ParentTheme      *Theme  `json:"parentTheme,omitempty"`

}

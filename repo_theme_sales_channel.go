package go_shopware_admin_sdk

import (
	"net/http"

)

type ThemeSalesChannelRepository struct {
	*GenericRepository[ThemeSalesChannel]
}

func NewThemeSalesChannelRepository(client *Client) *ThemeSalesChannelRepository {
	return &ThemeSalesChannelRepository{
		GenericRepository: NewGenericRepository[ThemeSalesChannel](client),
	}
}

func (t *ThemeSalesChannelRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[ThemeSalesChannel], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "theme-sales-channel")
}

func (t *ThemeSalesChannelRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[ThemeSalesChannel], *http.Response, error) {
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

func (t *ThemeSalesChannelRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "theme-sales-channel")
}

func (t *ThemeSalesChannelRepository) Upsert(ctx ApiContext, entity []ThemeSalesChannel) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "theme_sales_channel")
}

func (t *ThemeSalesChannelRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "theme_sales_channel")
}

type ThemeSalesChannel struct {

	SalesChannel      *SalesChannel  `json:"salesChannel,omitempty"`

	SalesChannelId      string  `json:"salesChannelId,omitempty"`

	Theme      *Theme  `json:"theme,omitempty"`

	ThemeId      string  `json:"themeId,omitempty"`

}

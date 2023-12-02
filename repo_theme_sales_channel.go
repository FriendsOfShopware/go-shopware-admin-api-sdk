package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type ThemeSalesChannelRepository ClientService

func (t ThemeSalesChannelRepository) Search(ctx ApiContext, criteria Criteria) (*ThemeSalesChannelCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/theme-sales-channel", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(ThemeSalesChannelCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ThemeSalesChannelRepository) SearchAll(ctx ApiContext, criteria Criteria) (*ThemeSalesChannelCollection, *http.Response, error) {
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

func (t ThemeSalesChannelRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/theme-sales-channel", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(SearchIdsResponse)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ThemeSalesChannelRepository) Upsert(ctx ApiContext, entity []ThemeSalesChannel) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"theme_sales_channel": {
		Entity:  "theme_sales_channel",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t ThemeSalesChannelRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"theme_sales_channel": {
		Entity:  "theme_sales_channel",
		Action:  "delete",
		Payload: payload,
	}})
}

type ThemeSalesChannel struct {

	SalesChannelId      string  `json:"salesChannelId,omitempty"`

	ThemeId      string  `json:"themeId,omitempty"`

	Theme      *Theme  `json:"theme,omitempty"`

	SalesChannel      *SalesChannel  `json:"salesChannel,omitempty"`

}

type ThemeSalesChannelCollection struct {
	EntityCollection

	Data []ThemeSalesChannel `json:"data"`
}

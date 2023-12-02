package go_shopware_admin_sdk

import (
	"net/http"

	"time"
)

type SalesChannelTypeRepository ClientService

func (t SalesChannelTypeRepository) Search(ctx ApiContext, criteria Criteria) (*SalesChannelTypeCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/sales-channel-type", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(SalesChannelTypeCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t SalesChannelTypeRepository) SearchAll(ctx ApiContext, criteria Criteria) (*SalesChannelTypeCollection, *http.Response, error) {
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

func (t SalesChannelTypeRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/sales-channel-type", criteria)

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

func (t SalesChannelTypeRepository) Upsert(ctx ApiContext, entity []SalesChannelType) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"sales_channel_type": {
		Entity:  "sales_channel_type",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t SalesChannelTypeRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"sales_channel_type": {
		Entity:  "sales_channel_type",
		Action:  "delete",
		Payload: payload,
	}})
}

type SalesChannelType struct {
	ScreenshotUrls interface{} `json:"screenshotUrls,omitempty"`

	Name string `json:"name,omitempty"`

	Manufacturer string `json:"manufacturer,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Translations []SalesChannelTypeTranslation `json:"translations,omitempty"`

	SalesChannels []SalesChannel `json:"salesChannels,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	DescriptionLong string `json:"descriptionLong,omitempty"`

	Translated interface{} `json:"translated,omitempty"`

	Id string `json:"id,omitempty"`

	CoverUrl string `json:"coverUrl,omitempty"`

	IconName string `json:"iconName,omitempty"`

	Description string `json:"description,omitempty"`
}

type SalesChannelTypeCollection struct {
	EntityCollection

	Data []SalesChannelType `json:"data"`
}

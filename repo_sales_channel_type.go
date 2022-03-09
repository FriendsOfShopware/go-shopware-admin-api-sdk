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
	Id string `json:"id,omitempty"`

	Description string `json:"description,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	CoverUrl string `json:"coverUrl,omitempty"`

	Manufacturer string `json:"manufacturer,omitempty"`

	Translations []SalesChannelTypeTranslation `json:"translations,omitempty"`

	IconName string `json:"iconName,omitempty"`

	DescriptionLong string `json:"descriptionLong,omitempty"`

	SalesChannels []SalesChannel `json:"salesChannels,omitempty"`

	Translated interface{} `json:"translated,omitempty"`

	ScreenshotUrls interface{} `json:"screenshotUrls,omitempty"`

	Name string `json:"name,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`
}

type SalesChannelTypeCollection struct {
	EntityCollection

	Data []SalesChannelType `json:"data"`
}

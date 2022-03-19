package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type SalesChannelTypeTranslationRepository ClientService

func (t SalesChannelTypeTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*SalesChannelTypeTranslationCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/sales-channel-type-translation", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(SalesChannelTypeTranslationCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t SalesChannelTypeTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*SalesChannelTypeTranslationCollection, *http.Response, error) {
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

func (t SalesChannelTypeTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/sales-channel-type-translation", criteria)

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

func (t SalesChannelTypeTranslationRepository) Upsert(ctx ApiContext, entity []SalesChannelTypeTranslation) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"sales_channel_type_translation": {
		Entity:  "sales_channel_type_translation",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t SalesChannelTypeTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"sales_channel_type_translation": {
		Entity:  "sales_channel_type_translation",
		Action:  "delete",
		Payload: payload,
	}})
}

type SalesChannelTypeTranslation struct {
	CreatedAt time.Time `json:"createdAt,omitempty"`

	SalesChannelTypeId string `json:"salesChannelTypeId,omitempty"`

	Name string `json:"name,omitempty"`

	DescriptionLong string `json:"descriptionLong,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	SalesChannelType *SalesChannelType `json:"salesChannelType,omitempty"`

	Language *Language `json:"language,omitempty"`

	Manufacturer string `json:"manufacturer,omitempty"`

	Description string `json:"description,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type SalesChannelTypeTranslationCollection struct {
	EntityCollection

	Data []SalesChannelTypeTranslation `json:"data"`
}

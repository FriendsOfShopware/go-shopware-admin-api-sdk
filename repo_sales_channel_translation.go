package go_shopware_admin_sdk

import (
	"net/http"

	"time"
)

type SalesChannelTranslationRepository ClientService

func (t SalesChannelTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*SalesChannelTranslationCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/sales-channel-translation", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(SalesChannelTranslationCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t SalesChannelTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*SalesChannelTranslationCollection, *http.Response, error) {
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

func (t SalesChannelTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/sales-channel-translation", criteria)

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

func (t SalesChannelTranslationRepository) Upsert(ctx ApiContext, entity []SalesChannelTranslation) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"sales_channel_translation": {
		Entity:  "sales_channel_translation",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t SalesChannelTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"sales_channel_translation": {
		Entity:  "sales_channel_translation",
		Action:  "delete",
		Payload: payload,
	}})
}

type SalesChannelTranslation struct {
	HomeMetaTitle string `json:"homeMetaTitle,omitempty"`

	HomeMetaDescription string `json:"homeMetaDescription,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Name string `json:"name,omitempty"`

	HomeEnabled bool `json:"homeEnabled,omitempty"`

	HomeName string `json:"homeName,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	SalesChannelId string `json:"salesChannelId,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	HomeSlotConfig interface{} `json:"homeSlotConfig,omitempty"`

	SalesChannel *SalesChannel `json:"salesChannel,omitempty"`

	HomeKeywords string `json:"homeKeywords,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	Language *Language `json:"language,omitempty"`
}

type SalesChannelTranslationCollection struct {
	EntityCollection

	Data []SalesChannelTranslation `json:"data"`
}

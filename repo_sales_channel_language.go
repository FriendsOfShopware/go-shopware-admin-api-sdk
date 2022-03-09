package go_shopware_admin_sdk

import (
	"net/http"
)

type SalesChannelLanguageRepository ClientService

func (t SalesChannelLanguageRepository) Search(ctx ApiContext, criteria Criteria) (*SalesChannelLanguageCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/sales-channel-language", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(SalesChannelLanguageCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t SalesChannelLanguageRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/sales-channel-language", criteria)

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

func (t SalesChannelLanguageRepository) Upsert(ctx ApiContext, entity []SalesChannelLanguage) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"sales_channel_language": {
		Entity:  "sales_channel_language",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t SalesChannelLanguageRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"sales_channel_language": {
		Entity:  "sales_channel_language",
		Action:  "delete",
		Payload: payload,
	}})
}

type SalesChannelLanguage struct {
	SalesChannel *SalesChannel `json:"salesChannel,omitempty"`

	Language *Language `json:"language,omitempty"`

	SalesChannelId string `json:"salesChannelId,omitempty"`

	LanguageId string `json:"languageId,omitempty"`
}

type SalesChannelLanguageCollection struct {
	EntityCollection

	Data []SalesChannelLanguage `json:"data"`
}

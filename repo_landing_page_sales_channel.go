package go_shopware_admin_sdk

import (
	"net/http"
)

type LandingPageSalesChannelRepository ClientService

func (t LandingPageSalesChannelRepository) Search(ctx ApiContext, criteria Criteria) (*LandingPageSalesChannelCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/landing-page-sales-channel", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(LandingPageSalesChannelCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t LandingPageSalesChannelRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/landing-page-sales-channel", criteria)

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

func (t LandingPageSalesChannelRepository) Upsert(ctx ApiContext, entity []LandingPageSalesChannel) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"landing_page_sales_channel": {
		Entity:  "landing_page_sales_channel",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t LandingPageSalesChannelRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"landing_page_sales_channel": {
		Entity:  "landing_page_sales_channel",
		Action:  "delete",
		Payload: payload,
	}})
}

type LandingPageSalesChannel struct {
	LandingPageId string `json:"landingPageId,omitempty"`

	LandingPageVersionId string `json:"landingPageVersionId,omitempty"`

	SalesChannelId string `json:"salesChannelId,omitempty"`

	LandingPage *LandingPage `json:"landingPage,omitempty"`

	SalesChannel *SalesChannel `json:"salesChannel,omitempty"`
}

type LandingPageSalesChannelCollection struct {
	EntityCollection

	Data []LandingPageSalesChannel `json:"data"`
}

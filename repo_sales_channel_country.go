package go_shopware_admin_sdk

import (
	"net/http"
)

type SalesChannelCountryRepository ClientService

func (t SalesChannelCountryRepository) Search(ctx ApiContext, criteria Criteria) (*SalesChannelCountryCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/sales-channel-country", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(SalesChannelCountryCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t SalesChannelCountryRepository) SearchAll(ctx ApiContext, criteria Criteria) (*SalesChannelCountryCollection, *http.Response, error) {
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

func (t SalesChannelCountryRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/sales-channel-country", criteria)

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

func (t SalesChannelCountryRepository) Upsert(ctx ApiContext, entity []SalesChannelCountry) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"sales_channel_country": {
		Entity:  "sales_channel_country",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t SalesChannelCountryRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"sales_channel_country": {
		Entity:  "sales_channel_country",
		Action:  "delete",
		Payload: payload,
	}})
}

type SalesChannelCountry struct {
	SalesChannelId string `json:"salesChannelId,omitempty"`

	CountryId string `json:"countryId,omitempty"`

	SalesChannel *SalesChannel `json:"salesChannel,omitempty"`

	Country *Country `json:"country,omitempty"`
}

type SalesChannelCountryCollection struct {
	EntityCollection

	Data []SalesChannelCountry `json:"data"`
}

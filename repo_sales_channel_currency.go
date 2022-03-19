package go_shopware_admin_sdk

import (
	"net/http"
)

type SalesChannelCurrencyRepository ClientService

func (t SalesChannelCurrencyRepository) Search(ctx ApiContext, criteria Criteria) (*SalesChannelCurrencyCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/sales-channel-currency", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(SalesChannelCurrencyCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t SalesChannelCurrencyRepository) SearchAll(ctx ApiContext, criteria Criteria) (*SalesChannelCurrencyCollection, *http.Response, error) {
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

func (t SalesChannelCurrencyRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/sales-channel-currency", criteria)

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

func (t SalesChannelCurrencyRepository) Upsert(ctx ApiContext, entity []SalesChannelCurrency) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"sales_channel_currency": {
		Entity:  "sales_channel_currency",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t SalesChannelCurrencyRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"sales_channel_currency": {
		Entity:  "sales_channel_currency",
		Action:  "delete",
		Payload: payload,
	}})
}

type SalesChannelCurrency struct {
	CurrencyId string `json:"currencyId,omitempty"`

	SalesChannel *SalesChannel `json:"salesChannel,omitempty"`

	Currency *Currency `json:"currency,omitempty"`

	SalesChannelId string `json:"salesChannelId,omitempty"`
}

type SalesChannelCurrencyCollection struct {
	EntityCollection

	Data []SalesChannelCurrency `json:"data"`
}

package go_shopware_admin_sdk

import (
	"net/http"

	"time"
)

type CurrencyCountryRoundingRepository ClientService

func (t CurrencyCountryRoundingRepository) Search(ctx ApiContext, criteria Criteria) (*CurrencyCountryRoundingCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/currency-country-rounding", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(CurrencyCountryRoundingCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t CurrencyCountryRoundingRepository) SearchAll(ctx ApiContext, criteria Criteria) (*CurrencyCountryRoundingCollection, *http.Response, error) {
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

func (t CurrencyCountryRoundingRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/currency-country-rounding", criteria)

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

func (t CurrencyCountryRoundingRepository) Upsert(ctx ApiContext, entity []CurrencyCountryRounding) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"currency_country_rounding": {
		Entity:  "currency_country_rounding",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t CurrencyCountryRoundingRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"currency_country_rounding": {
		Entity:  "currency_country_rounding",
		Action:  "delete",
		Payload: payload,
	}})
}

type CurrencyCountryRounding struct {
	CurrencyId string `json:"currencyId,omitempty"`

	Currency *Currency `json:"currency,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Id string `json:"id,omitempty"`

	CountryId string `json:"countryId,omitempty"`

	ItemRounding interface{} `json:"itemRounding,omitempty"`

	TotalRounding interface{} `json:"totalRounding,omitempty"`

	Country *Country `json:"country,omitempty"`
}

type CurrencyCountryRoundingCollection struct {
	EntityCollection

	Data []CurrencyCountryRounding `json:"data"`
}

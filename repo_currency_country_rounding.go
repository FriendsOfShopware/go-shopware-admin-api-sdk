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
	ItemRounding interface{} `json:"itemRounding,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Currency *Currency `json:"currency,omitempty"`

	Country *Country `json:"country,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Id string `json:"id,omitempty"`

	CurrencyId string `json:"currencyId,omitempty"`

	CountryId string `json:"countryId,omitempty"`

	TotalRounding interface{} `json:"totalRounding,omitempty"`
}

type CurrencyCountryRoundingCollection struct {
	EntityCollection

	Data []CurrencyCountryRounding `json:"data"`
}

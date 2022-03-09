package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type CurrencyRepository ClientService

func (t CurrencyRepository) Search(ctx ApiContext, criteria Criteria) (*CurrencyCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/currency", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(CurrencyCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t CurrencyRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/currency", criteria)

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

func (t CurrencyRepository) Upsert(ctx ApiContext, entity []Currency) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"currency": {
		Entity:  "currency",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t CurrencyRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"currency": {
		Entity:  "currency",
		Action:  "delete",
		Payload: payload,
	}})
}

type Currency struct {
	ProductExports []ProductExport `json:"productExports,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Translated interface{} `json:"translated,omitempty"`

	IsoCode string `json:"isoCode,omitempty"`

	Orders []Order `json:"orders,omitempty"`

	PromotionDiscountPrices []PromotionDiscountPrices `json:"promotionDiscountPrices,omitempty"`

	Translations []CurrencyTranslation `json:"translations,omitempty"`

	SalesChannelDefaultAssignments []SalesChannel `json:"salesChannelDefaultAssignments,omitempty"`

	SalesChannelDomains []SalesChannelDomain `json:"salesChannelDomains,omitempty"`

	TotalRounding interface{} `json:"totalRounding,omitempty"`

	Id string `json:"id,omitempty"`

	Factor float64 `json:"factor,omitempty"`

	Name string `json:"name,omitempty"`

	CountryRoundings []CurrencyCountryRounding `json:"countryRoundings,omitempty"`

	Symbol string `json:"symbol,omitempty"`

	IsSystemDefault bool `json:"isSystemDefault,omitempty"`

	ItemRounding interface{} `json:"itemRounding,omitempty"`

	SalesChannels []SalesChannel `json:"salesChannels,omitempty"`

	TaxFreeFrom float64 `json:"taxFreeFrom,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	ShortName string `json:"shortName,omitempty"`

	Position float64 `json:"position,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`
}

type CurrencyCollection struct {
	EntityCollection

	Data []Currency `json:"data"`
}

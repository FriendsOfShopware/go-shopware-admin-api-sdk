package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type CountryRepository ClientService

func (t CountryRepository) Search(ctx ApiContext, criteria Criteria) (*CountryCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/country", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(CountryCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t CountryRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/country", criteria)

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

func (t CountryRepository) Upsert(ctx ApiContext, entity []Country) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"country": {
		Entity:  "country",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t CountryRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"country": {
		Entity:  "country",
		Action:  "delete",
		Payload: payload,
	}})
}

type Country struct {
	CreatedAt time.Time `json:"createdAt,omitempty"`

	Name string `json:"name,omitempty"`

	ShippingAvailable bool `json:"shippingAvailable,omitempty"`

	DisplayStateInRegistration bool `json:"displayStateInRegistration,omitempty"`

	CustomerTax interface{} `json:"customerTax,omitempty"`

	CompanyTax interface{} `json:"companyTax,omitempty"`

	Translations []CountryTranslation `json:"translations,omitempty"`

	OrderAddresses []OrderAddress `json:"orderAddresses,omitempty"`

	Iso3 string `json:"iso3,omitempty"`

	ForceStateInRegistration bool `json:"forceStateInRegistration,omitempty"`

	Id string `json:"id,omitempty"`

	VatIdRequired bool `json:"vatIdRequired,omitempty"`

	CustomerAddresses []CustomerAddress `json:"customerAddresses,omitempty"`

	TaxFree bool `json:"taxFree,omitempty"`

	Active bool `json:"active,omitempty"`

	CompanyTaxFree bool `json:"companyTaxFree,omitempty"`

	SalesChannelDefaultAssignments []SalesChannel `json:"salesChannelDefaultAssignments,omitempty"`

	CurrencyCountryRoundings []CurrencyCountryRounding `json:"currencyCountryRoundings,omitempty"`

	SalesChannels []SalesChannel `json:"salesChannels,omitempty"`

	TaxRules []TaxRule `json:"taxRules,omitempty"`

	Translated interface{} `json:"translated,omitempty"`

	Position float64 `json:"position,omitempty"`

	CheckVatIdPattern bool `json:"checkVatIdPattern,omitempty"`

	VatIdPattern string `json:"vatIdPattern,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Iso string `json:"iso,omitempty"`

	States []CountryState `json:"states,omitempty"`
}

type CountryCollection struct {
	EntityCollection

	Data []Country `json:"data"`
}

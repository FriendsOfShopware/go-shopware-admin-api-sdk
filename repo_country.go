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

func (t CountryRepository) SearchAll(ctx ApiContext, criteria Criteria) (*CountryCollection, *http.Response, error) {
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

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	Name      string  `json:"name,omitempty"`

	Active      bool  `json:"active,omitempty"`

	CheckAdvancedPostalCodePattern      bool  `json:"checkAdvancedPostalCodePattern,omitempty"`

	CurrencyCountryRoundings      []CurrencyCountryRounding  `json:"currencyCountryRoundings,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Id      string  `json:"id,omitempty"`

	Iso3      string  `json:"iso3,omitempty"`

	ForceStateInRegistration      bool  `json:"forceStateInRegistration,omitempty"`

	VatIdRequired      bool  `json:"vatIdRequired,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	OrderAddresses      []OrderAddress  `json:"orderAddresses,omitempty"`

	ShippingAvailable      bool  `json:"shippingAvailable,omitempty"`

	PostalCodeRequired      bool  `json:"postalCodeRequired,omitempty"`

	Translations      []CountryTranslation  `json:"translations,omitempty"`

	Iso      string  `json:"iso,omitempty"`

	CompanyTax      interface{}  `json:"companyTax,omitempty"`

	States      []CountryState  `json:"states,omitempty"`

	SalesChannels      []SalesChannel  `json:"salesChannels,omitempty"`

	CustomerAddresses      []CustomerAddress  `json:"customerAddresses,omitempty"`

	TaxRules      []TaxRule  `json:"taxRules,omitempty"`

	CheckPostalCodePattern      bool  `json:"checkPostalCodePattern,omitempty"`

	AdvancedPostalCodePattern      string  `json:"advancedPostalCodePattern,omitempty"`

	AddressFormat      interface{}  `json:"addressFormat,omitempty"`

	Position      float64  `json:"position,omitempty"`

	DisplayStateInRegistration      bool  `json:"displayStateInRegistration,omitempty"`

	CheckVatIdPattern      bool  `json:"checkVatIdPattern,omitempty"`

	VatIdPattern      string  `json:"vatIdPattern,omitempty"`

	CustomerTax      interface{}  `json:"customerTax,omitempty"`

	DefaultPostalCodePattern      string  `json:"defaultPostalCodePattern,omitempty"`

	SalesChannelDefaultAssignments      []SalesChannel  `json:"salesChannelDefaultAssignments,omitempty"`

}

type CountryCollection struct {
	EntityCollection

	Data []Country `json:"data"`
}

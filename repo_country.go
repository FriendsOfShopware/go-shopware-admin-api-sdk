package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type CountryRepository struct {
	*GenericRepository[Country]
}

func NewCountryRepository(client *Client) *CountryRepository {
	return &CountryRepository{
		GenericRepository: NewGenericRepository[Country](client),
	}
}

func (t *CountryRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[Country], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "country")
}

func (t *CountryRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[Country], *http.Response, error) {
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

func (t *CountryRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "country")
}

func (t *CountryRepository) Upsert(ctx ApiContext, entity []Country) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "country")
}

func (t *CountryRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "country")
}

type Country struct {

	DisplayStateInRegistration      bool  `json:"displayStateInRegistration,omitempty"`

	States      []CountryState  `json:"states,omitempty"`

	OrderAddresses      []OrderAddress  `json:"orderAddresses,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	Position      float64  `json:"position,omitempty"`

	AddressFormat      interface{}  `json:"addressFormat,omitempty"`

	IsEu      bool  `json:"isEu,omitempty"`

	SalesChannelDefaultAssignments      []SalesChannel  `json:"salesChannelDefaultAssignments,omitempty"`

	Iso      string  `json:"iso,omitempty"`

	ForceStateInRegistration      bool  `json:"forceStateInRegistration,omitempty"`

	CheckPostalCodePattern      bool  `json:"checkPostalCodePattern,omitempty"`

	TaxRules      []TaxRule  `json:"taxRules,omitempty"`

	CompanyTax      interface{}  `json:"companyTax,omitempty"`

	Translations      []CountryTranslation  `json:"translations,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	CustomerTax      interface{}  `json:"customerTax,omitempty"`

	Name      string  `json:"name,omitempty"`

	Active      bool  `json:"active,omitempty"`

	DefaultPostalCodePattern      string  `json:"defaultPostalCodePattern,omitempty"`

	CurrencyCountryRoundings      []CurrencyCountryRounding  `json:"currencyCountryRoundings,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	CheckAdvancedPostalCodePattern      bool  `json:"checkAdvancedPostalCodePattern,omitempty"`

	CustomerAddresses      []CustomerAddress  `json:"customerAddresses,omitempty"`

	SalesChannels      []SalesChannel  `json:"salesChannels,omitempty"`

	VatIdPattern      string  `json:"vatIdPattern,omitempty"`

	ShippingAvailable      bool  `json:"shippingAvailable,omitempty"`

	Iso3      string  `json:"iso3,omitempty"`

	CheckVatIdPattern      bool  `json:"checkVatIdPattern,omitempty"`

	PostalCodeRequired      bool  `json:"postalCodeRequired,omitempty"`

	Id      string  `json:"id,omitempty"`

	AdvancedPostalCodePattern      string  `json:"advancedPostalCodePattern,omitempty"`

	VatIdRequired      bool  `json:"vatIdRequired,omitempty"`

}

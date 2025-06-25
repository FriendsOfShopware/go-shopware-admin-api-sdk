package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type CurrencyCountryRoundingRepository struct {
	*GenericRepository[CurrencyCountryRounding]
}

func NewCurrencyCountryRoundingRepository(client *Client) *CurrencyCountryRoundingRepository {
	return &CurrencyCountryRoundingRepository{
		GenericRepository: NewGenericRepository[CurrencyCountryRounding](client),
	}
}

func (t *CurrencyCountryRoundingRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[CurrencyCountryRounding], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "currency-country-rounding")
}

func (t *CurrencyCountryRoundingRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[CurrencyCountryRounding], *http.Response, error) {
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

func (t *CurrencyCountryRoundingRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "currency-country-rounding")
}

func (t *CurrencyCountryRoundingRepository) Upsert(ctx ApiContext, entity []CurrencyCountryRounding) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "currency_country_rounding")
}

func (t *CurrencyCountryRoundingRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "currency_country_rounding")
}

type CurrencyCountryRounding struct {

	Country      *Country  `json:"country,omitempty"`

	CountryId      string  `json:"countryId,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Currency      *Currency  `json:"currency,omitempty"`

	CurrencyId      string  `json:"currencyId,omitempty"`

	Id      string  `json:"id,omitempty"`

	ItemRounding      interface{}  `json:"itemRounding,omitempty"`

	TotalRounding      interface{}  `json:"totalRounding,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}

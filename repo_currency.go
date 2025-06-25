package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type CurrencyRepository struct {
	*GenericRepository[Currency]
}

func NewCurrencyRepository(client *Client) *CurrencyRepository {
	return &CurrencyRepository{
		GenericRepository: NewGenericRepository[Currency](client),
	}
}

func (t *CurrencyRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[Currency], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "currency")
}

func (t *CurrencyRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[Currency], *http.Response, error) {
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

func (t *CurrencyRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "currency")
}

func (t *CurrencyRepository) Upsert(ctx ApiContext, entity []Currency) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "currency")
}

func (t *CurrencyRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "currency")
}

type Currency struct {

	Symbol      string  `json:"symbol,omitempty"`

	ShortName      string  `json:"shortName,omitempty"`

	IsSystemDefault      bool  `json:"isSystemDefault,omitempty"`

	SalesChannelDefaultAssignments      []SalesChannel  `json:"salesChannelDefaultAssignments,omitempty"`

	Orders      []Order  `json:"orders,omitempty"`

	SalesChannelDomains      []SalesChannelDomain  `json:"salesChannelDomains,omitempty"`

	ItemRounding      interface{}  `json:"itemRounding,omitempty"`

	TotalRounding      interface{}  `json:"totalRounding,omitempty"`

	Name      string  `json:"name,omitempty"`

	Position      float64  `json:"position,omitempty"`

	TaxFreeFrom      float64  `json:"taxFreeFrom,omitempty"`

	CountryRoundings      []CurrencyCountryRounding  `json:"countryRoundings,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Factor      float64  `json:"factor,omitempty"`

	IsoCode      string  `json:"isoCode,omitempty"`

	Translations      []CurrencyTranslation  `json:"translations,omitempty"`

	SalesChannels      []SalesChannel  `json:"salesChannels,omitempty"`

	PromotionDiscountPrices      []PromotionDiscountPrices  `json:"promotionDiscountPrices,omitempty"`

	ProductExports      []ProductExport  `json:"productExports,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	Id      string  `json:"id,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

}

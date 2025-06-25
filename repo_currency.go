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

	CountryRoundings      []CurrencyCountryRounding  `json:"countryRoundings,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Factor      float64  `json:"factor,omitempty"`

	Id      string  `json:"id,omitempty"`

	IsSystemDefault      bool  `json:"isSystemDefault,omitempty"`

	IsoCode      string  `json:"isoCode,omitempty"`

	ItemRounding      interface{}  `json:"itemRounding,omitempty"`

	Name      string  `json:"name,omitempty"`

	Orders      []Order  `json:"orders,omitempty"`

	Position      float64  `json:"position,omitempty"`

	ProductExports      []ProductExport  `json:"productExports,omitempty"`

	PromotionDiscountPrices      []PromotionDiscountPrices  `json:"promotionDiscountPrices,omitempty"`

	SalesChannelDefaultAssignments      []SalesChannel  `json:"salesChannelDefaultAssignments,omitempty"`

	SalesChannelDomains      []SalesChannelDomain  `json:"salesChannelDomains,omitempty"`

	SalesChannels      []SalesChannel  `json:"salesChannels,omitempty"`

	ShortName      string  `json:"shortName,omitempty"`

	Symbol      string  `json:"symbol,omitempty"`

	TaxFreeFrom      float64  `json:"taxFreeFrom,omitempty"`

	TotalRounding      interface{}  `json:"totalRounding,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	Translations      []CurrencyTranslation  `json:"translations,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}

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

func (t CurrencyRepository) SearchAll(ctx ApiContext, criteria Criteria) (*CurrencyCollection, *http.Response, error) {
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
	CountryRoundings []CurrencyCountryRounding `json:"countryRoundings,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Name string `json:"name,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	Orders []Order `json:"orders,omitempty"`

	SalesChannels []SalesChannel `json:"salesChannels,omitempty"`

	TotalRounding interface{} `json:"totalRounding,omitempty"`

	Symbol string `json:"symbol,omitempty"`

	IsSystemDefault bool `json:"isSystemDefault,omitempty"`

	ItemRounding interface{} `json:"itemRounding,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	PromotionDiscountPrices []PromotionDiscountPrices `json:"promotionDiscountPrices,omitempty"`

	Id string `json:"id,omitempty"`

	Factor float64 `json:"factor,omitempty"`

	Position float64 `json:"position,omitempty"`

	TaxFreeFrom float64 `json:"taxFreeFrom,omitempty"`

	Translations []CurrencyTranslation `json:"translations,omitempty"`

	Translated interface{} `json:"translated,omitempty"`

	IsoCode string `json:"isoCode,omitempty"`

	ShortName string `json:"shortName,omitempty"`

	SalesChannelDefaultAssignments []SalesChannel `json:"salesChannelDefaultAssignments,omitempty"`

	SalesChannelDomains []SalesChannelDomain `json:"salesChannelDomains,omitempty"`

	ProductExports []ProductExport `json:"productExports,omitempty"`
}

type CurrencyCollection struct {
	EntityCollection

	Data []Currency `json:"data"`
}

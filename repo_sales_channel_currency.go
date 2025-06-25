package go_shopware_admin_sdk

import (
	"net/http"

)

type SalesChannelCurrencyRepository struct {
	*GenericRepository[SalesChannelCurrency]
}

func NewSalesChannelCurrencyRepository(client *Client) *SalesChannelCurrencyRepository {
	return &SalesChannelCurrencyRepository{
		GenericRepository: NewGenericRepository[SalesChannelCurrency](client),
	}
}

func (t *SalesChannelCurrencyRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[SalesChannelCurrency], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "sales-channel-currency")
}

func (t *SalesChannelCurrencyRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[SalesChannelCurrency], *http.Response, error) {
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

func (t *SalesChannelCurrencyRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "sales-channel-currency")
}

func (t *SalesChannelCurrencyRepository) Upsert(ctx ApiContext, entity []SalesChannelCurrency) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "sales_channel_currency")
}

func (t *SalesChannelCurrencyRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "sales_channel_currency")
}

type SalesChannelCurrency struct {

	Currency      *Currency  `json:"currency,omitempty"`

	CurrencyId      string  `json:"currencyId,omitempty"`

	SalesChannel      *SalesChannel  `json:"salesChannel,omitempty"`

	SalesChannelId      string  `json:"salesChannelId,omitempty"`

}

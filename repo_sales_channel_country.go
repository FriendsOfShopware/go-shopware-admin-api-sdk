package go_shopware_admin_sdk

import (
	"net/http"

)

type SalesChannelCountryRepository struct {
	*GenericRepository[SalesChannelCountry]
}

func NewSalesChannelCountryRepository(client *Client) *SalesChannelCountryRepository {
	return &SalesChannelCountryRepository{
		GenericRepository: NewGenericRepository[SalesChannelCountry](client),
	}
}

func (t *SalesChannelCountryRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[SalesChannelCountry], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "sales-channel-country")
}

func (t *SalesChannelCountryRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[SalesChannelCountry], *http.Response, error) {
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

func (t *SalesChannelCountryRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "sales-channel-country")
}

func (t *SalesChannelCountryRepository) Upsert(ctx ApiContext, entity []SalesChannelCountry) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "sales_channel_country")
}

func (t *SalesChannelCountryRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "sales_channel_country")
}

type SalesChannelCountry struct {

	CountryId      string  `json:"countryId,omitempty"`

	SalesChannel      *SalesChannel  `json:"salesChannel,omitempty"`

	Country      *Country  `json:"country,omitempty"`

	SalesChannelId      string  `json:"salesChannelId,omitempty"`

}

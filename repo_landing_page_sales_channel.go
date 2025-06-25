package go_shopware_admin_sdk

import (
	"net/http"

)

type LandingPageSalesChannelRepository struct {
	*GenericRepository[LandingPageSalesChannel]
}

func NewLandingPageSalesChannelRepository(client *Client) *LandingPageSalesChannelRepository {
	return &LandingPageSalesChannelRepository{
		GenericRepository: NewGenericRepository[LandingPageSalesChannel](client),
	}
}

func (t *LandingPageSalesChannelRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[LandingPageSalesChannel], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "landing-page-sales-channel")
}

func (t *LandingPageSalesChannelRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[LandingPageSalesChannel], *http.Response, error) {
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

func (t *LandingPageSalesChannelRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "landing-page-sales-channel")
}

func (t *LandingPageSalesChannelRepository) Upsert(ctx ApiContext, entity []LandingPageSalesChannel) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "landing_page_sales_channel")
}

func (t *LandingPageSalesChannelRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "landing_page_sales_channel")
}

type LandingPageSalesChannel struct {

	LandingPage      *LandingPage  `json:"landingPage,omitempty"`

	LandingPageId      string  `json:"landingPageId,omitempty"`

	LandingPageVersionId      string  `json:"landingPageVersionId,omitempty"`

	SalesChannel      *SalesChannel  `json:"salesChannel,omitempty"`

	SalesChannelId      string  `json:"salesChannelId,omitempty"`

}

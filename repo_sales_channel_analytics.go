package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type SalesChannelAnalyticsRepository struct {
	*GenericRepository[SalesChannelAnalytics]
}

func NewSalesChannelAnalyticsRepository(client *Client) *SalesChannelAnalyticsRepository {
	return &SalesChannelAnalyticsRepository{
		GenericRepository: NewGenericRepository[SalesChannelAnalytics](client),
	}
}

func (t *SalesChannelAnalyticsRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[SalesChannelAnalytics], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "sales-channel-analytics")
}

func (t *SalesChannelAnalyticsRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[SalesChannelAnalytics], *http.Response, error) {
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

func (t *SalesChannelAnalyticsRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "sales-channel-analytics")
}

func (t *SalesChannelAnalyticsRepository) Upsert(ctx ApiContext, entity []SalesChannelAnalytics) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "sales_channel_analytics")
}

func (t *SalesChannelAnalyticsRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "sales_channel_analytics")
}

type SalesChannelAnalytics struct {

	Active      bool  `json:"active,omitempty"`

	AnonymizeIp      bool  `json:"anonymizeIp,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Id      string  `json:"id,omitempty"`

	SalesChannel      *SalesChannel  `json:"salesChannel,omitempty"`

	TrackOrders      bool  `json:"trackOrders,omitempty"`

	TrackingId      string  `json:"trackingId,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}

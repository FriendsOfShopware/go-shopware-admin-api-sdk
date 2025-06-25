package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type NumberRangeSalesChannelRepository struct {
	*GenericRepository[NumberRangeSalesChannel]
}

func NewNumberRangeSalesChannelRepository(client *Client) *NumberRangeSalesChannelRepository {
	return &NumberRangeSalesChannelRepository{
		GenericRepository: NewGenericRepository[NumberRangeSalesChannel](client),
	}
}

func (t *NumberRangeSalesChannelRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[NumberRangeSalesChannel], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "number-range-sales-channel")
}

func (t *NumberRangeSalesChannelRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[NumberRangeSalesChannel], *http.Response, error) {
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

func (t *NumberRangeSalesChannelRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "number-range-sales-channel")
}

func (t *NumberRangeSalesChannelRepository) Upsert(ctx ApiContext, entity []NumberRangeSalesChannel) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "number_range_sales_channel")
}

func (t *NumberRangeSalesChannelRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "number_range_sales_channel")
}

type NumberRangeSalesChannel struct {

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Id      string  `json:"id,omitempty"`

	NumberRange      *NumberRange  `json:"numberRange,omitempty"`

	NumberRangeId      string  `json:"numberRangeId,omitempty"`

	NumberRangeType      *NumberRangeType  `json:"numberRangeType,omitempty"`

	NumberRangeTypeId      string  `json:"numberRangeTypeId,omitempty"`

	SalesChannel      *SalesChannel  `json:"salesChannel,omitempty"`

	SalesChannelId      string  `json:"salesChannelId,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}

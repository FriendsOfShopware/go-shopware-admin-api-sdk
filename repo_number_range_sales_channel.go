package go_shopware_admin_sdk

import (
	"net/http"

	"time"
)

type NumberRangeSalesChannelRepository ClientService

func (t NumberRangeSalesChannelRepository) Search(ctx ApiContext, criteria Criteria) (*NumberRangeSalesChannelCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/number-range-sales-channel", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(NumberRangeSalesChannelCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t NumberRangeSalesChannelRepository) SearchAll(ctx ApiContext, criteria Criteria) (*NumberRangeSalesChannelCollection, *http.Response, error) {
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

func (t NumberRangeSalesChannelRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/number-range-sales-channel", criteria)

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

func (t NumberRangeSalesChannelRepository) Upsert(ctx ApiContext, entity []NumberRangeSalesChannel) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"number_range_sales_channel": {
		Entity:  "number_range_sales_channel",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t NumberRangeSalesChannelRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"number_range_sales_channel": {
		Entity:  "number_range_sales_channel",
		Action:  "delete",
		Payload: payload,
	}})
}

type NumberRangeSalesChannel struct {
	NumberRangeTypeId string `json:"numberRangeTypeId,omitempty"`

	SalesChannel *SalesChannel `json:"salesChannel,omitempty"`

	NumberRangeType *NumberRangeType `json:"numberRangeType,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Id string `json:"id,omitempty"`

	NumberRangeId string `json:"numberRangeId,omitempty"`

	SalesChannelId string `json:"salesChannelId,omitempty"`

	NumberRange *NumberRange `json:"numberRange,omitempty"`
}

type NumberRangeSalesChannelCollection struct {
	EntityCollection

	Data []NumberRangeSalesChannel `json:"data"`
}

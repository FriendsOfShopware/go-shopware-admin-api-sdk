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
	NumberRange *NumberRange `json:"numberRange,omitempty"`

	SalesChannel *SalesChannel `json:"salesChannel,omitempty"`

	NumberRangeType *NumberRangeType `json:"numberRangeType,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Id string `json:"id,omitempty"`

	NumberRangeId string `json:"numberRangeId,omitempty"`

	SalesChannelId string `json:"salesChannelId,omitempty"`

	NumberRangeTypeId string `json:"numberRangeTypeId,omitempty"`
}

type NumberRangeSalesChannelCollection struct {
	EntityCollection

	Data []NumberRangeSalesChannel `json:"data"`
}

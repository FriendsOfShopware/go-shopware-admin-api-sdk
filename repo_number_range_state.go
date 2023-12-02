package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type NumberRangeStateRepository ClientService

func (t NumberRangeStateRepository) Search(ctx ApiContext, criteria Criteria) (*NumberRangeStateCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/number-range-state", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(NumberRangeStateCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t NumberRangeStateRepository) SearchAll(ctx ApiContext, criteria Criteria) (*NumberRangeStateCollection, *http.Response, error) {
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

func (t NumberRangeStateRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/number-range-state", criteria)

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

func (t NumberRangeStateRepository) Upsert(ctx ApiContext, entity []NumberRangeState) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"number_range_state": {
		Entity:  "number_range_state",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t NumberRangeStateRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"number_range_state": {
		Entity:  "number_range_state",
		Action:  "delete",
		Payload: payload,
	}})
}

type NumberRangeState struct {

	LastValue      float64  `json:"lastValue,omitempty"`

	NumberRange      *NumberRange  `json:"numberRange,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Id      string  `json:"id,omitempty"`

	NumberRangeId      string  `json:"numberRangeId,omitempty"`

}

type NumberRangeStateCollection struct {
	EntityCollection

	Data []NumberRangeState `json:"data"`
}

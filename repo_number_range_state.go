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
	Id string `json:"id,omitempty"`

	NumberRangeId string `json:"numberRangeId,omitempty"`

	LastValue float64 `json:"lastValue,omitempty"`

	NumberRange *NumberRange `json:"numberRange,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type NumberRangeStateCollection struct {
	EntityCollection

	Data []NumberRangeState `json:"data"`
}

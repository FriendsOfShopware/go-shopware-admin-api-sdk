package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type NumberRangeRepository ClientService

func (t NumberRangeRepository) Search(ctx ApiContext, criteria Criteria) (*NumberRangeCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/number-range", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(NumberRangeCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t NumberRangeRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/number-range", criteria)

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

func (t NumberRangeRepository) Upsert(ctx ApiContext, entity []NumberRange) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"number_range": {
		Entity:  "number_range",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t NumberRangeRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"number_range": {
		Entity:  "number_range",
		Action:  "delete",
		Payload: payload,
	}})
}

type NumberRange struct {
	CustomFields interface{} `json:"customFields,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Id string `json:"id,omitempty"`

	Global bool `json:"global,omitempty"`

	Start float64 `json:"start,omitempty"`

	Name string `json:"name,omitempty"`

	Translations []NumberRangeTranslation `json:"translations,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Translated interface{} `json:"translated,omitempty"`

	Description string `json:"description,omitempty"`

	Type *NumberRangeType `json:"type,omitempty"`

	NumberRangeSalesChannels []NumberRangeSalesChannel `json:"numberRangeSalesChannels,omitempty"`

	TypeId string `json:"typeId,omitempty"`

	Pattern string `json:"pattern,omitempty"`

	State *NumberRangeState `json:"state,omitempty"`
}

type NumberRangeCollection struct {
	EntityCollection

	Data []NumberRange `json:"data"`
}

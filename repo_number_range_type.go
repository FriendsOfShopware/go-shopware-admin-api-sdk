package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type NumberRangeTypeRepository ClientService

func (t NumberRangeTypeRepository) Search(ctx ApiContext, criteria Criteria) (*NumberRangeTypeCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/number-range-type", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(NumberRangeTypeCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t NumberRangeTypeRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/number-range-type", criteria)

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

func (t NumberRangeTypeRepository) Upsert(ctx ApiContext, entity []NumberRangeType) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"number_range_type": {
		Entity:  "number_range_type",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t NumberRangeTypeRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"number_range_type": {
		Entity:  "number_range_type",
		Action:  "delete",
		Payload: payload,
	}})
}

type NumberRangeType struct {
	TypeName string `json:"typeName,omitempty"`

	NumberRanges []NumberRange `json:"numberRanges,omitempty"`

	Translations []NumberRangeTypeTranslation `json:"translations,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Id string `json:"id,omitempty"`

	TechnicalName string `json:"technicalName,omitempty"`

	Global bool `json:"global,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	NumberRangeSalesChannels []NumberRangeSalesChannel `json:"numberRangeSalesChannels,omitempty"`

	Translated interface{} `json:"translated,omitempty"`
}

type NumberRangeTypeCollection struct {
	EntityCollection

	Data []NumberRangeType `json:"data"`
}

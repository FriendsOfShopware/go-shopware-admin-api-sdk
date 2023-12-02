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

func (t NumberRangeTypeRepository) SearchAll(ctx ApiContext, criteria Criteria) (*NumberRangeTypeCollection, *http.Response, error) {
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

	Translations []NumberRangeTypeTranslation `json:"translations,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Translated interface{} `json:"translated,omitempty"`

	Id string `json:"id,omitempty"`

	TechnicalName string `json:"technicalName,omitempty"`

	NumberRanges []NumberRange `json:"numberRanges,omitempty"`

	NumberRangeSalesChannels []NumberRangeSalesChannel `json:"numberRangeSalesChannels,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Global bool `json:"global,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`
}

type NumberRangeTypeCollection struct {
	EntityCollection

	Data []NumberRangeType `json:"data"`
}

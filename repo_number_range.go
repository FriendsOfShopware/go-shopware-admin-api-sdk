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

func (t NumberRangeRepository) SearchAll(ctx ApiContext, criteria Criteria) (*NumberRangeCollection, *http.Response, error) {
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

	TypeId      string  `json:"typeId,omitempty"`

	Pattern      string  `json:"pattern,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	State      *NumberRangeState  `json:"state,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	Global      bool  `json:"global,omitempty"`

	Name      string  `json:"name,omitempty"`

	Description      string  `json:"description,omitempty"`

	Start      float64  `json:"start,omitempty"`

	Type      *NumberRangeType  `json:"type,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Id      string  `json:"id,omitempty"`

	NumberRangeSalesChannels      []NumberRangeSalesChannel  `json:"numberRangeSalesChannels,omitempty"`

	Translations      []NumberRangeTranslation  `json:"translations,omitempty"`

}

type NumberRangeCollection struct {
	EntityCollection

	Data []NumberRange `json:"data"`
}

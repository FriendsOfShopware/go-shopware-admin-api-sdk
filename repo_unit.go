package go_shopware_admin_sdk

import (
	"net/http"

	"time"
)

type UnitRepository ClientService

func (t UnitRepository) Search(ctx ApiContext, criteria Criteria) (*UnitCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/unit", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(UnitCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t UnitRepository) SearchAll(ctx ApiContext, criteria Criteria) (*UnitCollection, *http.Response, error) {
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

func (t UnitRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/unit", criteria)

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

func (t UnitRepository) Upsert(ctx ApiContext, entity []Unit) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"unit": {
		Entity:  "unit",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t UnitRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"unit": {
		Entity:  "unit",
		Action:  "delete",
		Payload: payload,
	}})
}

type Unit struct {
	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	Products []Product `json:"products,omitempty"`

	Translations []UnitTranslation `json:"translations,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	ShortCode string `json:"shortCode,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Translated interface{} `json:"translated,omitempty"`
}

type UnitCollection struct {
	EntityCollection

	Data []Unit `json:"data"`
}

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
	CreatedAt time.Time `json:"createdAt,omitempty"`

	Translated interface{} `json:"translated,omitempty"`

	Id string `json:"id,omitempty"`

	ShortCode string `json:"shortCode,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	Products []Product `json:"products,omitempty"`

	Name string `json:"name,omitempty"`

	Translations []UnitTranslation `json:"translations,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type UnitCollection struct {
	EntityCollection

	Data []Unit `json:"data"`
}

package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type CountryStateRepository ClientService

func (t CountryStateRepository) Search(ctx ApiContext, criteria Criteria) (*CountryStateCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/country-state", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(CountryStateCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t CountryStateRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/country-state", criteria)

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

func (t CountryStateRepository) Upsert(ctx ApiContext, entity []CountryState) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"country_state": {
		Entity:  "country_state",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t CountryStateRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"country_state": {
		Entity:  "country_state",
		Action:  "delete",
		Payload: payload,
	}})
}

type CountryState struct {
	Id string `json:"id,omitempty"`

	Position float64 `json:"position,omitempty"`

	CountryId string `json:"countryId,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	OrderAddresses []OrderAddress `json:"orderAddresses,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	ShortCode string `json:"shortCode,omitempty"`

	Translations []CountryStateTranslation `json:"translations,omitempty"`

	CustomerAddresses []CustomerAddress `json:"customerAddresses,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Translated interface{} `json:"translated,omitempty"`

	Name string `json:"name,omitempty"`

	Active bool `json:"active,omitempty"`

	Country *Country `json:"country,omitempty"`
}

type CountryStateCollection struct {
	EntityCollection

	Data []CountryState `json:"data"`
}

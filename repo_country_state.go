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

func (t CountryStateRepository) SearchAll(ctx ApiContext, criteria Criteria) (*CountryStateCollection, *http.Response, error) {
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

	Id      string  `json:"id,omitempty"`

	CountryId      string  `json:"countryId,omitempty"`

	ShortCode      string  `json:"shortCode,omitempty"`

	Name      string  `json:"name,omitempty"`

	Position      float64  `json:"position,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Active      bool  `json:"active,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	Country      *Country  `json:"country,omitempty"`

	Translations      []CountryStateTranslation  `json:"translations,omitempty"`

	CustomerAddresses      []CustomerAddress  `json:"customerAddresses,omitempty"`

	OrderAddresses      []OrderAddress  `json:"orderAddresses,omitempty"`

}

type CountryStateCollection struct {
	EntityCollection

	Data []CountryState `json:"data"`
}

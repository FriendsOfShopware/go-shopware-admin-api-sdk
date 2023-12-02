package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type CountryStateTranslationRepository ClientService

func (t CountryStateTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*CountryStateTranslationCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/country-state-translation", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(CountryStateTranslationCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t CountryStateTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*CountryStateTranslationCollection, *http.Response, error) {
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

func (t CountryStateTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/country-state-translation", criteria)

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

func (t CountryStateTranslationRepository) Upsert(ctx ApiContext, entity []CountryStateTranslation) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"country_state_translation": {
		Entity:  "country_state_translation",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t CountryStateTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"country_state_translation": {
		Entity:  "country_state_translation",
		Action:  "delete",
		Payload: payload,
	}})
}

type CountryStateTranslation struct {

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	CountryStateId      string  `json:"countryStateId,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	CountryState      *CountryState  `json:"countryState,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	Name      string  `json:"name,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

}

type CountryStateTranslationCollection struct {
	EntityCollection

	Data []CountryStateTranslation `json:"data"`
}

package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type CountryTranslationRepository ClientService

func (t CountryTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*CountryTranslationCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/country-translation", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(CountryTranslationCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t CountryTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/country-translation", criteria)

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

func (t CountryTranslationRepository) Upsert(ctx ApiContext, entity []CountryTranslation) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"country_translation": {
		Entity:  "country_translation",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t CountryTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"country_translation": {
		Entity:  "country_translation",
		Action:  "delete",
		Payload: payload,
	}})
}

type CountryTranslation struct {
	CustomFields interface{} `json:"customFields,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	CountryId string `json:"countryId,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	Country *Country `json:"country,omitempty"`

	Language *Language `json:"language,omitempty"`

	Name string `json:"name,omitempty"`
}

type CountryTranslationCollection struct {
	EntityCollection

	Data []CountryTranslation `json:"data"`
}

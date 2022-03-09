package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type LocaleRepository ClientService

func (t LocaleRepository) Search(ctx ApiContext, criteria Criteria) (*LocaleCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/locale", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(LocaleCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t LocaleRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/locale", criteria)

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

func (t LocaleRepository) Upsert(ctx ApiContext, entity []Locale) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"locale": {
		Entity:  "locale",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t LocaleRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"locale": {
		Entity:  "locale",
		Action:  "delete",
		Payload: payload,
	}})
}

type Locale struct {
	CustomFields interface{} `json:"customFields,omitempty"`

	Translations []LocaleTranslation `json:"translations,omitempty"`

	Translated interface{} `json:"translated,omitempty"`

	Code string `json:"code,omitempty"`

	Name string `json:"name,omitempty"`

	Territory string `json:"territory,omitempty"`

	Languages []Language `json:"languages,omitempty"`

	Users []User `json:"users,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Id string `json:"id,omitempty"`
}

type LocaleCollection struct {
	EntityCollection

	Data []Locale `json:"data"`
}

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

func (t LocaleRepository) SearchAll(ctx ApiContext, criteria Criteria) (*LocaleCollection, *http.Response, error) {
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
	Languages []Language `json:"languages,omitempty"`

	Users []User `json:"users,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Translated interface{} `json:"translated,omitempty"`

	Name string `json:"name,omitempty"`

	Code string `json:"code,omitempty"`

	Territory string `json:"territory,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	Translations []LocaleTranslation `json:"translations,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Id string `json:"id,omitempty"`
}

type LocaleCollection struct {
	EntityCollection

	Data []Locale `json:"data"`
}

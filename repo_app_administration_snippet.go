package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type AppAdministrationSnippetRepository ClientService

func (t AppAdministrationSnippetRepository) Search(ctx ApiContext, criteria Criteria) (*AppAdministrationSnippetCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/app-administration-snippet", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(AppAdministrationSnippetCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t AppAdministrationSnippetRepository) SearchAll(ctx ApiContext, criteria Criteria) (*AppAdministrationSnippetCollection, *http.Response, error) {
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

func (t AppAdministrationSnippetRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/app-administration-snippet", criteria)

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

func (t AppAdministrationSnippetRepository) Upsert(ctx ApiContext, entity []AppAdministrationSnippet) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"app_administration_snippet": {
		Entity:  "app_administration_snippet",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t AppAdministrationSnippetRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"app_administration_snippet": {
		Entity:  "app_administration_snippet",
		Action:  "delete",
		Payload: payload,
	}})
}

type AppAdministrationSnippet struct {

	LocaleId      string  `json:"localeId,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Id      string  `json:"id,omitempty"`

	Value      string  `json:"value,omitempty"`

	AppId      string  `json:"appId,omitempty"`

}

type AppAdministrationSnippetCollection struct {
	EntityCollection

	Data []AppAdministrationSnippet `json:"data"`
}

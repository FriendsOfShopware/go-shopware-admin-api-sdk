package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type SnippetRepository ClientService

func (t SnippetRepository) Search(ctx ApiContext, criteria Criteria) (*SnippetCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/snippet", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(SnippetCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t SnippetRepository) SearchAll(ctx ApiContext, criteria Criteria) (*SnippetCollection, *http.Response, error) {
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

func (t SnippetRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/snippet", criteria)

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

func (t SnippetRepository) Upsert(ctx ApiContext, entity []Snippet) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"snippet": {
		Entity:  "snippet",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t SnippetRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"snippet": {
		Entity:  "snippet",
		Action:  "delete",
		Payload: payload,
	}})
}

type Snippet struct {
	SetId string `json:"setId,omitempty"`

	TranslationKey string `json:"translationKey,omitempty"`

	Author string `json:"author,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Id string `json:"id,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	Set *SnippetSet `json:"set,omitempty"`

	Value string `json:"value,omitempty"`
}

type SnippetCollection struct {
	EntityCollection

	Data []Snippet `json:"data"`
}

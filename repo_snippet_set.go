package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type SnippetSetRepository ClientService

func (t SnippetSetRepository) Search(ctx ApiContext, criteria Criteria) (*SnippetSetCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/snippet-set", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(SnippetSetCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t SnippetSetRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/snippet-set", criteria)

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

func (t SnippetSetRepository) Upsert(ctx ApiContext, entity []SnippetSet) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"snippet_set": {
		Entity:  "snippet_set",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t SnippetSetRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"snippet_set": {
		Entity:  "snippet_set",
		Action:  "delete",
		Payload: payload,
	}})
}

type SnippetSet struct {
	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Iso string `json:"iso,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	BaseFile string `json:"baseFile,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	Snippets []Snippet `json:"snippets,omitempty"`

	SalesChannelDomains []SalesChannelDomain `json:"salesChannelDomains,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type SnippetSetCollection struct {
	EntityCollection

	Data []SnippetSet `json:"data"`
}

package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type VersionRepository ClientService

func (t VersionRepository) Search(ctx ApiContext, criteria Criteria) (*VersionCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/version", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(VersionCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t VersionRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/version", criteria)

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

func (t VersionRepository) Upsert(ctx ApiContext, entity []Version) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"version": {
		Entity:  "version",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t VersionRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"version": {
		Entity:  "version",
		Action:  "delete",
		Payload: payload,
	}})
}

type Version struct {
	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Commits []VersionCommit `json:"commits,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type VersionCollection struct {
	EntityCollection

	Data []Version `json:"data"`
}

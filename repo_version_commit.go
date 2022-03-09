package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type VersionCommitRepository ClientService

func (t VersionCommitRepository) Search(ctx ApiContext, criteria Criteria) (*VersionCommitCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/version-commit", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(VersionCommitCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t VersionCommitRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/version-commit", criteria)

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

func (t VersionCommitRepository) Upsert(ctx ApiContext, entity []VersionCommit) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"version_commit": {
		Entity:  "version_commit",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t VersionCommitRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"version_commit": {
		Entity:  "version_commit",
		Action:  "delete",
		Payload: payload,
	}})
}

type VersionCommit struct {
	AutoIncrement float64 `json:"autoIncrement,omitempty"`

	IsMerge bool `json:"isMerge,omitempty"`

	Data []VersionCommitData `json:"data,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Id string `json:"id,omitempty"`

	VersionId string `json:"versionId,omitempty"`

	IntegrationId string `json:"integrationId,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	UserId string `json:"userId,omitempty"`

	Message string `json:"message,omitempty"`

	Version *Version `json:"version,omitempty"`
}

type VersionCommitCollection struct {
	EntityCollection

	Data []VersionCommit `json:"data"`
}

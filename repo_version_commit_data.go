package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type VersionCommitDataRepository ClientService

func (t VersionCommitDataRepository) Search(ctx ApiContext, criteria Criteria) (*VersionCommitDataCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/version-commit-data", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(VersionCommitDataCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t VersionCommitDataRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/version-commit-data", criteria)

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

func (t VersionCommitDataRepository) Upsert(ctx ApiContext, entity []VersionCommitData) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"version_commit_data": {
		Entity:  "version_commit_data",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t VersionCommitDataRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"version_commit_data": {
		Entity:  "version_commit_data",
		Action:  "delete",
		Payload: payload,
	}})
}

type VersionCommitData struct {
	Payload interface{} `json:"payload,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Id string `json:"id,omitempty"`

	VersionCommitId string `json:"versionCommitId,omitempty"`

	AutoIncrement float64 `json:"autoIncrement,omitempty"`

	Action string `json:"action,omitempty"`

	EntityId interface{} `json:"entityId,omitempty"`

	Commit *VersionCommit `json:"commit,omitempty"`

	UserId string `json:"userId,omitempty"`

	IntegrationId string `json:"integrationId,omitempty"`

	EntityName string `json:"entityName,omitempty"`
}

type VersionCommitDataCollection struct {
	EntityCollection

	Data []VersionCommitData `json:"data"`
}

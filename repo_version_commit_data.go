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

func (t VersionCommitDataRepository) SearchAll(ctx ApiContext, criteria Criteria) (*VersionCommitDataCollection, *http.Response, error) {
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
	Id string `json:"id,omitempty"`

	VersionCommitId string `json:"versionCommitId,omitempty"`

	Commit *VersionCommit `json:"commit,omitempty"`

	UserId string `json:"userId,omitempty"`

	AutoIncrement float64 `json:"autoIncrement,omitempty"`

	EntityName string `json:"entityName,omitempty"`

	EntityId interface{} `json:"entityId,omitempty"`

	Action string `json:"action,omitempty"`

	Payload interface{} `json:"payload,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	IntegrationId string `json:"integrationId,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`
}

type VersionCommitDataCollection struct {
	EntityCollection

	Data []VersionCommitData `json:"data"`
}

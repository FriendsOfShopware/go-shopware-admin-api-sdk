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

func (t VersionCommitRepository) SearchAll(ctx ApiContext, criteria Criteria) (*VersionCommitCollection, *http.Response, error) {
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

	Version      *Version  `json:"version,omitempty"`

	Id      string  `json:"id,omitempty"`

	VersionId      string  `json:"versionId,omitempty"`

	UserId      string  `json:"userId,omitempty"`

	AutoIncrement      float64  `json:"autoIncrement,omitempty"`

	IsMerge      bool  `json:"isMerge,omitempty"`

	Data      []VersionCommitData  `json:"data,omitempty"`

	IntegrationId      string  `json:"integrationId,omitempty"`

	Message      string  `json:"message,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}

type VersionCommitCollection struct {
	EntityCollection

	Data []VersionCommit `json:"data"`
}

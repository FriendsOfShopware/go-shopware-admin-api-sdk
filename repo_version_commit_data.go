package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type VersionCommitDataRepository struct {
	*GenericRepository[VersionCommitData]
}

func NewVersionCommitDataRepository(client *Client) *VersionCommitDataRepository {
	return &VersionCommitDataRepository{
		GenericRepository: NewGenericRepository[VersionCommitData](client),
	}
}

func (t *VersionCommitDataRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[VersionCommitData], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "version-commit-data")
}

func (t *VersionCommitDataRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[VersionCommitData], *http.Response, error) {
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

func (t *VersionCommitDataRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "version-commit-data")
}

func (t *VersionCommitDataRepository) Upsert(ctx ApiContext, entity []VersionCommitData) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "version_commit_data")
}

func (t *VersionCommitDataRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "version_commit_data")
}

type VersionCommitData struct {

	Payload      interface{}  `json:"payload,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Id      string  `json:"id,omitempty"`

	VersionCommitId      string  `json:"versionCommitId,omitempty"`

	EntityName      string  `json:"entityName,omitempty"`

	EntityId      interface{}  `json:"entityId,omitempty"`

	Action      string  `json:"action,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Commit      *VersionCommit  `json:"commit,omitempty"`

	UserId      string  `json:"userId,omitempty"`

	IntegrationId      string  `json:"integrationId,omitempty"`

	AutoIncrement      float64  `json:"autoIncrement,omitempty"`

}

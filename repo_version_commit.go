package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type VersionCommitRepository struct {
	*GenericRepository[VersionCommit]
}

func NewVersionCommitRepository(client *Client) *VersionCommitRepository {
	return &VersionCommitRepository{
		GenericRepository: NewGenericRepository[VersionCommit](client),
	}
}

func (t *VersionCommitRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[VersionCommit], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "version-commit")
}

func (t *VersionCommitRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[VersionCommit], *http.Response, error) {
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

func (t *VersionCommitRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "version-commit")
}

func (t *VersionCommitRepository) Upsert(ctx ApiContext, entity []VersionCommit) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "version_commit")
}

func (t *VersionCommitRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "version_commit")
}

type VersionCommit struct {

	IntegrationId      string  `json:"integrationId,omitempty"`

	AutoIncrement      float64  `json:"autoIncrement,omitempty"`

	Data      []VersionCommitData  `json:"data,omitempty"`

	Version      *Version  `json:"version,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Id      string  `json:"id,omitempty"`

	VersionId      string  `json:"versionId,omitempty"`

	UserId      string  `json:"userId,omitempty"`

	IsMerge      bool  `json:"isMerge,omitempty"`

	Message      string  `json:"message,omitempty"`

}

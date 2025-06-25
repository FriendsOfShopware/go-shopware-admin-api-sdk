package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type VersionRepository struct {
	*GenericRepository[Version]
}

func NewVersionRepository(client *Client) *VersionRepository {
	return &VersionRepository{
		GenericRepository: NewGenericRepository[Version](client),
	}
}

func (t *VersionRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[Version], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "version")
}

func (t *VersionRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[Version], *http.Response, error) {
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

func (t *VersionRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "version")
}

func (t *VersionRepository) Upsert(ctx ApiContext, entity []Version) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "version")
}

func (t *VersionRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "version")
}

type Version struct {

	Commits      []VersionCommit  `json:"commits,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Id      string  `json:"id,omitempty"`

	Name      string  `json:"name,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}

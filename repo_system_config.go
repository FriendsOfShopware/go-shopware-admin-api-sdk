package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type SystemConfigRepository struct {
	*GenericRepository[SystemConfig]
}

func NewSystemConfigRepository(client *Client) *SystemConfigRepository {
	return &SystemConfigRepository{
		GenericRepository: NewGenericRepository[SystemConfig](client),
	}
}

func (t *SystemConfigRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[SystemConfig], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "system-config")
}

func (t *SystemConfigRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[SystemConfig], *http.Response, error) {
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

func (t *SystemConfigRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "system-config")
}

func (t *SystemConfigRepository) Upsert(ctx ApiContext, entity []SystemConfig) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "system_config")
}

func (t *SystemConfigRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "system_config")
}

type SystemConfig struct {

	ConfigurationKey      string  `json:"configurationKey,omitempty"`

	ConfigurationValue      interface{}  `json:"configurationValue,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Id      string  `json:"id,omitempty"`

	SalesChannel      *SalesChannel  `json:"salesChannel,omitempty"`

	SalesChannelId      string  `json:"salesChannelId,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}

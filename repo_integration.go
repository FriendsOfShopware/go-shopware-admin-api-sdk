package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type IntegrationRepository struct {
	*GenericRepository[Integration]
}

func NewIntegrationRepository(client *Client) *IntegrationRepository {
	return &IntegrationRepository{
		GenericRepository: NewGenericRepository[Integration](client),
	}
}

func (t *IntegrationRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[Integration], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "integration")
}

func (t *IntegrationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[Integration], *http.Response, error) {
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

func (t *IntegrationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "integration")
}

func (t *IntegrationRepository) Upsert(ctx ApiContext, entity []Integration) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "integration")
}

func (t *IntegrationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "integration")
}

type Integration struct {

	Admin      bool  `json:"admin,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	DeletedAt      time.Time  `json:"deletedAt,omitempty"`

	AclRoles      []AclRole  `json:"aclRoles,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Id      string  `json:"id,omitempty"`

	Label      string  `json:"label,omitempty"`

	LastUsageAt      time.Time  `json:"lastUsageAt,omitempty"`

	App      *App  `json:"app,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	CreatedNotifications      []Notification  `json:"createdNotifications,omitempty"`

	AccessKey      string  `json:"accessKey,omitempty"`

	SecretAccessKey      interface{}  `json:"secretAccessKey,omitempty"`

}

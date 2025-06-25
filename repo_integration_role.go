package go_shopware_admin_sdk

import (
	"net/http"

)

type IntegrationRoleRepository struct {
	*GenericRepository[IntegrationRole]
}

func NewIntegrationRoleRepository(client *Client) *IntegrationRoleRepository {
	return &IntegrationRoleRepository{
		GenericRepository: NewGenericRepository[IntegrationRole](client),
	}
}

func (t *IntegrationRoleRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[IntegrationRole], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "integration-role")
}

func (t *IntegrationRoleRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[IntegrationRole], *http.Response, error) {
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

func (t *IntegrationRoleRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "integration-role")
}

func (t *IntegrationRoleRepository) Upsert(ctx ApiContext, entity []IntegrationRole) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "integration_role")
}

func (t *IntegrationRoleRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "integration_role")
}

type IntegrationRole struct {

	IntegrationId      string  `json:"integrationId,omitempty"`

	AclRoleId      string  `json:"aclRoleId,omitempty"`

	Integration      *Integration  `json:"integration,omitempty"`

	Role      *AclRole  `json:"role,omitempty"`

}

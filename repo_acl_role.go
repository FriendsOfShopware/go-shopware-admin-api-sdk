package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type AclRoleRepository struct {
	*GenericRepository[AclRole]
}

func NewAclRoleRepository(client *Client) *AclRoleRepository {
	return &AclRoleRepository{
		GenericRepository: NewGenericRepository[AclRole](client),
	}
}

func (t *AclRoleRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[AclRole], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "acl-role")
}

func (t *AclRoleRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[AclRole], *http.Response, error) {
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

func (t *AclRoleRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "acl-role")
}

func (t *AclRoleRepository) Upsert(ctx ApiContext, entity []AclRole) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "acl_role")
}

func (t *AclRoleRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "acl_role")
}

type AclRole struct {

	Id      string  `json:"id,omitempty"`

	Users      []User  `json:"users,omitempty"`

	App      *App  `json:"app,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Name      string  `json:"name,omitempty"`

	Description      string  `json:"description,omitempty"`

	Privileges      interface{}  `json:"privileges,omitempty"`

	DeletedAt      time.Time  `json:"deletedAt,omitempty"`

	Integrations      []Integration  `json:"integrations,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

}

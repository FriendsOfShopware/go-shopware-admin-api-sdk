package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type AclUserRoleRepository struct {
	*GenericRepository[AclUserRole]
}

func NewAclUserRoleRepository(client *Client) *AclUserRoleRepository {
	return &AclUserRoleRepository{
		GenericRepository: NewGenericRepository[AclUserRole](client),
	}
}

func (t *AclUserRoleRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[AclUserRole], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "acl-user-role")
}

func (t *AclUserRoleRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[AclUserRole], *http.Response, error) {
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

func (t *AclUserRoleRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "acl-user-role")
}

func (t *AclUserRoleRepository) Upsert(ctx ApiContext, entity []AclUserRole) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "acl_user_role")
}

func (t *AclUserRoleRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "acl_user_role")
}

type AclUserRole struct {

	UserId      string  `json:"userId,omitempty"`

	AclRoleId      string  `json:"aclRoleId,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	User      *User  `json:"user,omitempty"`

	AclRole      *AclRole  `json:"aclRole,omitempty"`

}

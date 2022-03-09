package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type AclUserRoleRepository ClientService

func (t AclUserRoleRepository) Search(ctx ApiContext, criteria Criteria) (*AclUserRoleCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/acl-user-role", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(AclUserRoleCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t AclUserRoleRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/acl-user-role", criteria)

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

func (t AclUserRoleRepository) Upsert(ctx ApiContext, entity []AclUserRole) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"acl_user_role": {
		Entity:  "acl_user_role",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t AclUserRoleRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"acl_user_role": {
		Entity:  "acl_user_role",
		Action:  "delete",
		Payload: payload,
	}})
}

type AclUserRole struct {
	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	User *User `json:"user,omitempty"`

	AclRole *AclRole `json:"aclRole,omitempty"`

	UserId string `json:"userId,omitempty"`

	AclRoleId string `json:"aclRoleId,omitempty"`
}

type AclUserRoleCollection struct {
	EntityCollection

	Data []AclUserRole `json:"data"`
}

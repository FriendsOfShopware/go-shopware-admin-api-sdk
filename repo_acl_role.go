package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type AclRoleRepository ClientService

func (t AclRoleRepository) Search(ctx ApiContext, criteria Criteria) (*AclRoleCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/acl-role", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(AclRoleCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t AclRoleRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/acl-role", criteria)

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

func (t AclRoleRepository) Upsert(ctx ApiContext, entity []AclRole) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"acl_role": {
		Entity:  "acl_role",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t AclRoleRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"acl_role": {
		Entity:  "acl_role",
		Action:  "delete",
		Payload: payload,
	}})
}

type AclRole struct {
	Id string `json:"id,omitempty"`

	Description string `json:"description,omitempty"`

	Privileges interface{} `json:"privileges,omitempty"`

	DeletedAt time.Time `json:"deletedAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Name string `json:"name,omitempty"`

	Users []User `json:"users,omitempty"`

	App *App `json:"app,omitempty"`

	Integrations []Integration `json:"integrations,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`
}

type AclRoleCollection struct {
	EntityCollection

	Data []AclRole `json:"data"`
}

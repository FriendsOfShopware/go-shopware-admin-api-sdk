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

func (t AclRoleRepository) SearchAll(ctx ApiContext, criteria Criteria) (*AclRoleCollection, *http.Response, error) {
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
	Users []User `json:"users,omitempty"`

	App *App `json:"app,omitempty"`

	Integrations []Integration `json:"integrations,omitempty"`

	Id string `json:"id,omitempty"`

	Description string `json:"description,omitempty"`

	Privileges interface{} `json:"privileges,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Name string `json:"name,omitempty"`

	DeletedAt time.Time `json:"deletedAt,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`
}

type AclRoleCollection struct {
	EntityCollection

	Data []AclRole `json:"data"`
}

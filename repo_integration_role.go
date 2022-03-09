package go_shopware_admin_sdk

import (
	"net/http"
)

type IntegrationRoleRepository ClientService

func (t IntegrationRoleRepository) Search(ctx ApiContext, criteria Criteria) (*IntegrationRoleCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/integration-role", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(IntegrationRoleCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t IntegrationRoleRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/integration-role", criteria)

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

func (t IntegrationRoleRepository) Upsert(ctx ApiContext, entity []IntegrationRole) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"integration_role": {
		Entity:  "integration_role",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t IntegrationRoleRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"integration_role": {
		Entity:  "integration_role",
		Action:  "delete",
		Payload: payload,
	}})
}

type IntegrationRole struct {
	IntegrationId string `json:"integrationId,omitempty"`

	AclRoleId string `json:"aclRoleId,omitempty"`

	Integration *Integration `json:"integration,omitempty"`

	Role *AclRole `json:"role,omitempty"`
}

type IntegrationRoleCollection struct {
	EntityCollection

	Data []IntegrationRole `json:"data"`
}

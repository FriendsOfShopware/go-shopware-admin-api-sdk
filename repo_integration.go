package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type IntegrationRepository ClientService

func (t IntegrationRepository) Search(ctx ApiContext, criteria Criteria) (*IntegrationCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/integration", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(IntegrationCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t IntegrationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/integration", criteria)

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

func (t IntegrationRepository) Upsert(ctx ApiContext, entity []Integration) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"integration": {
		Entity:  "integration",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t IntegrationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"integration": {
		Entity:  "integration",
		Action:  "delete",
		Payload: payload,
	}})
}

type Integration struct {
	Label string `json:"label,omitempty"`

	WriteAccess bool `json:"writeAccess,omitempty"`

	LastUsageAt time.Time `json:"lastUsageAt,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Admin bool `json:"admin,omitempty"`

	AclRoles []AclRole `json:"aclRoles,omitempty"`

	Id string `json:"id,omitempty"`

	AccessKey string `json:"accessKey,omitempty"`

	DeletedAt time.Time `json:"deletedAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	SecretAccessKey interface{} `json:"secretAccessKey,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	App *App `json:"app,omitempty"`

	CreatedNotifications []Notification `json:"createdNotifications,omitempty"`
}

type IntegrationCollection struct {
	EntityCollection

	Data []Integration `json:"data"`
}

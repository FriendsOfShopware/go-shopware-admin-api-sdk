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

func (t IntegrationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*IntegrationCollection, *http.Response, error) {
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
	LastUsageAt time.Time `json:"lastUsageAt,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Id string `json:"id,omitempty"`

	Label string `json:"label,omitempty"`

	AccessKey string `json:"accessKey,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	CreatedNotifications []Notification `json:"createdNotifications,omitempty"`

	SecretAccessKey interface{} `json:"secretAccessKey,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	Admin bool `json:"admin,omitempty"`

	DeletedAt time.Time `json:"deletedAt,omitempty"`

	App *App `json:"app,omitempty"`

	AclRoles []AclRole `json:"aclRoles,omitempty"`
}

type IntegrationCollection struct {
	EntityCollection

	Data []Integration `json:"data"`
}

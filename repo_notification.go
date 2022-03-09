package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type NotificationRepository ClientService

func (t NotificationRepository) Search(ctx ApiContext, criteria Criteria) (*NotificationCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/notification", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(NotificationCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t NotificationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/notification", criteria)

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

func (t NotificationRepository) Upsert(ctx ApiContext, entity []Notification) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"notification": {
		Entity:  "notification",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t NotificationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"notification": {
		Entity:  "notification",
		Action:  "delete",
		Payload: payload,
	}})
}

type Notification struct {
	CreatedByIntegrationId string `json:"createdByIntegrationId,omitempty"`

	CreatedByUserId string `json:"createdByUserId,omitempty"`

	Id string `json:"id,omitempty"`

	Status string `json:"status,omitempty"`

	AdminOnly bool `json:"adminOnly,omitempty"`

	CreatedByUser *User `json:"createdByUser,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Message string `json:"message,omitempty"`

	RequiredPrivileges interface{} `json:"requiredPrivileges,omitempty"`

	CreatedByIntegration *Integration `json:"createdByIntegration,omitempty"`
}

type NotificationCollection struct {
	EntityCollection

	Data []Notification `json:"data"`
}

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

func (t NotificationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*NotificationCollection, *http.Response, error) {
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
	CreatedByUser *User `json:"createdByUser,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Status string `json:"status,omitempty"`

	Message string `json:"message,omitempty"`

	AdminOnly bool `json:"adminOnly,omitempty"`

	CreatedByIntegrationId string `json:"createdByIntegrationId,omitempty"`

	CreatedByIntegration *Integration `json:"createdByIntegration,omitempty"`

	Id string `json:"id,omitempty"`

	RequiredPrivileges interface{} `json:"requiredPrivileges,omitempty"`

	CreatedByUserId string `json:"createdByUserId,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type NotificationCollection struct {
	EntityCollection

	Data []Notification `json:"data"`
}

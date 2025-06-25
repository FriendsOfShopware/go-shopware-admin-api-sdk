package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type NotificationRepository struct {
	*GenericRepository[Notification]
}

func NewNotificationRepository(client *Client) *NotificationRepository {
	return &NotificationRepository{
		GenericRepository: NewGenericRepository[Notification](client),
	}
}

func (t *NotificationRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[Notification], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "notification")
}

func (t *NotificationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[Notification], *http.Response, error) {
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

func (t *NotificationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "notification")
}

func (t *NotificationRepository) Upsert(ctx ApiContext, entity []Notification) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "notification")
}

func (t *NotificationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "notification")
}

type Notification struct {
	Id             string    `json:"id,omitempty"`
	Status         string    `json:"status,omitempty"`
	Message        string    `json:"message,omitempty"`
	AdminOnly      bool      `json:"adminOnly,omitempty"`
	RequiredPrivileges []string `json:"requiredPrivileges,omitempty"`
	CreatedByIntegration *Integration `json:"createdByIntegration,omitempty"`
	CreatedByUser *User `json:"createdByUser,omitempty"`
	CreatedAt      time.Time `json:"createdAt,omitempty"`
	UpdatedAt      time.Time `json:"updatedAt,omitempty"`
}

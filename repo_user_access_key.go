package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type UserAccessKeyRepository struct {
	*GenericRepository[UserAccessKey]
}

func NewUserAccessKeyRepository(client *Client) *UserAccessKeyRepository {
	return &UserAccessKeyRepository{
		GenericRepository: NewGenericRepository[UserAccessKey](client),
	}
}

func (t *UserAccessKeyRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[UserAccessKey], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "user-access-key")
}

func (t *UserAccessKeyRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[UserAccessKey], *http.Response, error) {
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

func (t *UserAccessKeyRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "user-access-key")
}

func (t *UserAccessKeyRepository) Upsert(ctx ApiContext, entity []UserAccessKey) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "user_access_key")
}

func (t *UserAccessKeyRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "user_access_key")
}

type UserAccessKey struct {

	LastUsageAt      time.Time  `json:"lastUsageAt,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Id      string  `json:"id,omitempty"`

	UserId      string  `json:"userId,omitempty"`

	AccessKey      string  `json:"accessKey,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	User      *User  `json:"user,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	SecretAccessKey      interface{}  `json:"secretAccessKey,omitempty"`

}

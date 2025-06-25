package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type UserConfigRepository struct {
	*GenericRepository[UserConfig]
}

func NewUserConfigRepository(client *Client) *UserConfigRepository {
	return &UserConfigRepository{
		GenericRepository: NewGenericRepository[UserConfig](client),
	}
}

func (t *UserConfigRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[UserConfig], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "user-config")
}

func (t *UserConfigRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[UserConfig], *http.Response, error) {
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

func (t *UserConfigRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "user-config")
}

func (t *UserConfigRepository) Upsert(ctx ApiContext, entity []UserConfig) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "user_config")
}

func (t *UserConfigRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "user_config")
}

type UserConfig struct {

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Id      string  `json:"id,omitempty"`

	UserId      string  `json:"userId,omitempty"`

	Key      string  `json:"key,omitempty"`

	Value      interface{}  `json:"value,omitempty"`

	User      *User  `json:"user,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

}

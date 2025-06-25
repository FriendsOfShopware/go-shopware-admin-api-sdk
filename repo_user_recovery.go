package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type UserRecoveryRepository struct {
	*GenericRepository[UserRecovery]
}

func NewUserRecoveryRepository(client *Client) *UserRecoveryRepository {
	return &UserRecoveryRepository{
		GenericRepository: NewGenericRepository[UserRecovery](client),
	}
}

func (t *UserRecoveryRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[UserRecovery], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "user-recovery")
}

func (t *UserRecoveryRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[UserRecovery], *http.Response, error) {
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

func (t *UserRecoveryRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "user-recovery")
}

func (t *UserRecoveryRepository) Upsert(ctx ApiContext, entity []UserRecovery) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "user_recovery")
}

func (t *UserRecoveryRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "user_recovery")
}

type UserRecovery struct {

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Hash      string  `json:"hash,omitempty"`

	Id      string  `json:"id,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	User      *User  `json:"user,omitempty"`

	UserId      string  `json:"userId,omitempty"`

}

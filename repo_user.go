package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type UserRepository struct {
	*GenericRepository[User]
}

func NewUserRepository(client *Client) *UserRepository {
	return &UserRepository{
		GenericRepository: NewGenericRepository[User](client),
	}
}

func (t *UserRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[User], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "user")
}

func (t *UserRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[User], *http.Response, error) {
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

func (t *UserRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "user")
}

func (t *UserRepository) Upsert(ctx ApiContext, entity []User) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "user")
}

func (t *UserRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "user")
}

type User struct {

	AccessKeys      []UserAccessKey  `json:"accessKeys,omitempty"`

	AclRoles      []AclRole  `json:"aclRoles,omitempty"`

	Active      bool  `json:"active,omitempty"`

	Admin      bool  `json:"admin,omitempty"`

	AvatarId      string  `json:"avatarId,omitempty"`

	AvatarMedia      *Media  `json:"avatarMedia,omitempty"`

	Configs      []UserConfig  `json:"configs,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CreatedCustomers      []Customer  `json:"createdCustomers,omitempty"`

	CreatedNotifications      []Notification  `json:"createdNotifications,omitempty"`

	CreatedOrders      []Order  `json:"createdOrders,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Email      string  `json:"email,omitempty"`

	FirstName      string  `json:"firstName,omitempty"`

	Id      string  `json:"id,omitempty"`

	ImportExportLogEntries      []ImportExportLog  `json:"importExportLogEntries,omitempty"`

	LastName      string  `json:"lastName,omitempty"`

	LastUpdatedPasswordAt      time.Time  `json:"lastUpdatedPasswordAt,omitempty"`

	Locale      *Locale  `json:"locale,omitempty"`

	LocaleId      string  `json:"localeId,omitempty"`

	Media      []Media  `json:"media,omitempty"`

	Password      interface{}  `json:"password,omitempty"`

	RecoveryUser      *UserRecovery  `json:"recoveryUser,omitempty"`

	StateMachineHistoryEntries      []StateMachineHistory  `json:"stateMachineHistoryEntries,omitempty"`

	StoreToken      string  `json:"storeToken,omitempty"`

	TimeZone      string  `json:"timeZone,omitempty"`

	Title      string  `json:"title,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	UpdatedCustomers      []Customer  `json:"updatedCustomers,omitempty"`

	UpdatedOrders      []Order  `json:"updatedOrders,omitempty"`

	Username      string  `json:"username,omitempty"`

}

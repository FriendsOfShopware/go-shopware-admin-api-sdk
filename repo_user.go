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

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CreatedNotifications      []Notification  `json:"createdNotifications,omitempty"`

	Locale      *Locale  `json:"locale,omitempty"`

	ImportExportLogEntries      []ImportExportLog  `json:"importExportLogEntries,omitempty"`

	Active      bool  `json:"active,omitempty"`

	LastUpdatedPasswordAt      time.Time  `json:"lastUpdatedPasswordAt,omitempty"`

	AvatarMedia      *Media  `json:"avatarMedia,omitempty"`

	AclRoles      []AclRole  `json:"aclRoles,omitempty"`

	RecoveryUser      *UserRecovery  `json:"recoveryUser,omitempty"`

	CreatedOrders      []Order  `json:"createdOrders,omitempty"`

	Password      interface{}  `json:"password,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	AccessKeys      []UserAccessKey  `json:"accessKeys,omitempty"`

	LastName      string  `json:"lastName,omitempty"`

	Media      []Media  `json:"media,omitempty"`

	StateMachineHistoryEntries      []StateMachineHistory  `json:"stateMachineHistoryEntries,omitempty"`

	Title      string  `json:"title,omitempty"`

	Admin      bool  `json:"admin,omitempty"`

	TimeZone      string  `json:"timeZone,omitempty"`

	Configs      []UserConfig  `json:"configs,omitempty"`

	UpdatedCustomers      []Customer  `json:"updatedCustomers,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	FirstName      string  `json:"firstName,omitempty"`

	Email      string  `json:"email,omitempty"`

	AvatarId      string  `json:"avatarId,omitempty"`

	StoreToken      string  `json:"storeToken,omitempty"`

	UpdatedOrders      []Order  `json:"updatedOrders,omitempty"`

	CreatedCustomers      []Customer  `json:"createdCustomers,omitempty"`

	Id      string  `json:"id,omitempty"`

	LocaleId      string  `json:"localeId,omitempty"`

	Username      string  `json:"username,omitempty"`

}

package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type UserRepository ClientService

func (t UserRepository) Search(ctx ApiContext, criteria Criteria) (*UserCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/user", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(UserCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t UserRepository) SearchAll(ctx ApiContext, criteria Criteria) (*UserCollection, *http.Response, error) {
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

func (t UserRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/user", criteria)

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

func (t UserRepository) Upsert(ctx ApiContext, entity []User) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"user": {
		Entity:  "user",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t UserRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"user": {
		Entity:  "user",
		Action:  "delete",
		Payload: payload,
	}})
}

type User struct {

	Email      string  `json:"email,omitempty"`

	Username      string  `json:"username,omitempty"`

	Password      interface{}  `json:"password,omitempty"`

	FirstName      string  `json:"firstName,omitempty"`

	Configs      []UserConfig  `json:"configs,omitempty"`

	StoreToken      string  `json:"storeToken,omitempty"`

	UpdatedOrders      []Order  `json:"updatedOrders,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Id      string  `json:"id,omitempty"`

	Title      string  `json:"title,omitempty"`

	Admin      bool  `json:"admin,omitempty"`

	Locale      *Locale  `json:"locale,omitempty"`

	AvatarId      string  `json:"avatarId,omitempty"`

	StateMachineHistoryEntries      []StateMachineHistory  `json:"stateMachineHistoryEntries,omitempty"`

	AvatarMedia      *Media  `json:"avatarMedia,omitempty"`

	AccessKeys      []UserAccessKey  `json:"accessKeys,omitempty"`

	ImportExportLogEntries      []ImportExportLog  `json:"importExportLogEntries,omitempty"`

	UpdatedCustomers      []Customer  `json:"updatedCustomers,omitempty"`

	CreatedNotifications      []Notification  `json:"createdNotifications,omitempty"`

	LocaleId      string  `json:"localeId,omitempty"`

	LastName      string  `json:"lastName,omitempty"`

	TimeZone      string  `json:"timeZone,omitempty"`

	AclRoles      []AclRole  `json:"aclRoles,omitempty"`

	CreatedOrders      []Order  `json:"createdOrders,omitempty"`

	CreatedCustomers      []Customer  `json:"createdCustomers,omitempty"`

	Active      bool  `json:"active,omitempty"`

	LastUpdatedPasswordAt      time.Time  `json:"lastUpdatedPasswordAt,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Media      []Media  `json:"media,omitempty"`

	RecoveryUser      *UserRecovery  `json:"recoveryUser,omitempty"`

}

type UserCollection struct {
	EntityCollection

	Data []User `json:"data"`
}

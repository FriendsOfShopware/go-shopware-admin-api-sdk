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
	Id string `json:"id,omitempty"`

	TimeZone string `json:"timeZone,omitempty"`

	AvatarMedia *Media `json:"avatarMedia,omitempty"`

	CreatedOrders []Order `json:"createdOrders,omitempty"`

	UpdatedOrders []Order `json:"updatedOrders,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Title string `json:"title,omitempty"`

	Email string `json:"email,omitempty"`

	ImportExportLogEntries []ImportExportLog `json:"importExportLogEntries,omitempty"`

	CreatedCustomers []Customer `json:"createdCustomers,omitempty"`

	Admin bool `json:"admin,omitempty"`

	LastUpdatedPasswordAt time.Time `json:"lastUpdatedPasswordAt,omitempty"`

	StoreToken string `json:"storeToken,omitempty"`

	LocaleId string `json:"localeId,omitempty"`

	Username string `json:"username,omitempty"`

	FirstName string `json:"firstName,omitempty"`

	Active bool `json:"active,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	AvatarId string `json:"avatarId,omitempty"`

	Media []Media `json:"media,omitempty"`

	AccessKeys []UserAccessKey `json:"accessKeys,omitempty"`

	RecoveryUser *UserRecovery `json:"recoveryUser,omitempty"`

	UpdatedCustomers []Customer `json:"updatedCustomers,omitempty"`

	CreatedNotifications []Notification `json:"createdNotifications,omitempty"`

	Password interface{} `json:"password,omitempty"`

	Configs []UserConfig `json:"configs,omitempty"`

	AclRoles []AclRole `json:"aclRoles,omitempty"`

	LastName string `json:"lastName,omitempty"`

	Locale *Locale `json:"locale,omitempty"`

	StateMachineHistoryEntries []StateMachineHistory `json:"stateMachineHistoryEntries,omitempty"`
}

type UserCollection struct {
	EntityCollection

	Data []User `json:"data"`
}

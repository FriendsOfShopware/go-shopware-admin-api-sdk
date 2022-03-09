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
	LocaleId string `json:"localeId,omitempty"`

	LastName string `json:"lastName,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	Locale *Locale `json:"locale,omitempty"`

	ImportExportLogEntries []ImportExportLog `json:"importExportLogEntries,omitempty"`

	AclRoles []AclRole `json:"aclRoles,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Active bool `json:"active,omitempty"`

	StateMachineHistoryEntries []StateMachineHistory `json:"stateMachineHistoryEntries,omitempty"`

	RecoveryUser *UserRecovery `json:"recoveryUser,omitempty"`

	CreatedNotifications []Notification `json:"createdNotifications,omitempty"`

	AvatarMedia *Media `json:"avatarMedia,omitempty"`

	AccessKeys []UserAccessKey `json:"accessKeys,omitempty"`

	Configs []UserConfig `json:"configs,omitempty"`

	Title string `json:"title,omitempty"`

	Email string `json:"email,omitempty"`

	Media []Media `json:"media,omitempty"`

	Username string `json:"username,omitempty"`

	LastUpdatedPasswordAt time.Time `json:"lastUpdatedPasswordAt,omitempty"`

	UpdatedOrders []Order `json:"updatedOrders,omitempty"`

	AvatarId string `json:"avatarId,omitempty"`

	Password interface{} `json:"password,omitempty"`

	FirstName string `json:"firstName,omitempty"`

	Admin bool `json:"admin,omitempty"`

	CreatedOrders []Order `json:"createdOrders,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	TimeZone string `json:"timeZone,omitempty"`

	StoreToken string `json:"storeToken,omitempty"`

	Id string `json:"id,omitempty"`
}

type UserCollection struct {
	EntityCollection

	Data []User `json:"data"`
}

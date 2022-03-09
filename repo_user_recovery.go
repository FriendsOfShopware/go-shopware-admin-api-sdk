package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type UserRecoveryRepository ClientService

func (t UserRecoveryRepository) Search(ctx ApiContext, criteria Criteria) (*UserRecoveryCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/user-recovery", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(UserRecoveryCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t UserRecoveryRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/user-recovery", criteria)

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

func (t UserRecoveryRepository) Upsert(ctx ApiContext, entity []UserRecovery) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"user_recovery": {
		Entity:  "user_recovery",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t UserRecoveryRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"user_recovery": {
		Entity:  "user_recovery",
		Action:  "delete",
		Payload: payload,
	}})
}

type UserRecovery struct {
	Id string `json:"id,omitempty"`

	Hash string `json:"hash,omitempty"`

	UserId string `json:"userId,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	User *User `json:"user,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type UserRecoveryCollection struct {
	EntityCollection

	Data []UserRecovery `json:"data"`
}

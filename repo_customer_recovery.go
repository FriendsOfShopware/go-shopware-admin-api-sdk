package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type CustomerRecoveryRepository ClientService

func (t CustomerRecoveryRepository) Search(ctx ApiContext, criteria Criteria) (*CustomerRecoveryCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/customer-recovery", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(CustomerRecoveryCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t CustomerRecoveryRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/customer-recovery", criteria)

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

func (t CustomerRecoveryRepository) Upsert(ctx ApiContext, entity []CustomerRecovery) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"customer_recovery": {
		Entity:  "customer_recovery",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t CustomerRecoveryRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"customer_recovery": {
		Entity:  "customer_recovery",
		Action:  "delete",
		Payload: payload,
	}})
}

type CustomerRecovery struct {
	Id string `json:"id,omitempty"`

	Hash string `json:"hash,omitempty"`

	CustomerId string `json:"customerId,omitempty"`

	Customer *Customer `json:"customer,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type CustomerRecoveryCollection struct {
	EntityCollection

	Data []CustomerRecovery `json:"data"`
}

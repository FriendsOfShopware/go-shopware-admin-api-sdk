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

func (t CustomerRecoveryRepository) SearchAll(ctx ApiContext, criteria Criteria) (*CustomerRecoveryCollection, *http.Response, error) {
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
	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Id string `json:"id,omitempty"`

	Hash string `json:"hash,omitempty"`

	CustomerId string `json:"customerId,omitempty"`

	Customer *Customer `json:"customer,omitempty"`
}

type CustomerRecoveryCollection struct {
	EntityCollection

	Data []CustomerRecovery `json:"data"`
}

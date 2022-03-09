package go_shopware_admin_sdk

import (
	"net/http"
)

type CustomerTagRepository ClientService

func (t CustomerTagRepository) Search(ctx ApiContext, criteria Criteria) (*CustomerTagCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/customer-tag", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(CustomerTagCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t CustomerTagRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/customer-tag", criteria)

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

func (t CustomerTagRepository) Upsert(ctx ApiContext, entity []CustomerTag) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"customer_tag": {
		Entity:  "customer_tag",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t CustomerTagRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"customer_tag": {
		Entity:  "customer_tag",
		Action:  "delete",
		Payload: payload,
	}})
}

type CustomerTag struct {
	CustomerId string `json:"customerId,omitempty"`

	TagId string `json:"tagId,omitempty"`

	Customer *Customer `json:"customer,omitempty"`

	Tag *Tag `json:"tag,omitempty"`
}

type CustomerTagCollection struct {
	EntityCollection

	Data []CustomerTag `json:"data"`
}

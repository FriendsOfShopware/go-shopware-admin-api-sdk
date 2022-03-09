package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type OrderCustomerRepository ClientService

func (t OrderCustomerRepository) Search(ctx ApiContext, criteria Criteria) (*OrderCustomerCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/order-customer", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(OrderCustomerCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t OrderCustomerRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/order-customer", criteria)

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

func (t OrderCustomerRepository) Upsert(ctx ApiContext, entity []OrderCustomer) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"order_customer": {
		Entity:  "order_customer",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t OrderCustomerRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"order_customer": {
		Entity:  "order_customer",
		Action:  "delete",
		Payload: payload,
	}})
}

type OrderCustomer struct {
	Id string `json:"id,omitempty"`

	LastName string `json:"lastName,omitempty"`

	Company string `json:"company,omitempty"`

	CustomerNumber string `json:"customerNumber,omitempty"`

	Order *Order `json:"order,omitempty"`

	VersionId string `json:"versionId,omitempty"`

	CustomerId string `json:"customerId,omitempty"`

	OrderId string `json:"orderId,omitempty"`

	Title string `json:"title,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	OrderVersionId string `json:"orderVersionId,omitempty"`

	VatIds interface{} `json:"vatIds,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	Customer *Customer `json:"customer,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Email string `json:"email,omitempty"`

	SalutationId string `json:"salutationId,omitempty"`

	FirstName string `json:"firstName,omitempty"`

	Salutation *Salutation `json:"salutation,omitempty"`

	RemoteAddress interface{} `json:"remoteAddress,omitempty"`
}

type OrderCustomerCollection struct {
	EntityCollection

	Data []OrderCustomer `json:"data"`
}

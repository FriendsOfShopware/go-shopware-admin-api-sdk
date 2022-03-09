package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type CustomerAddressRepository ClientService

func (t CustomerAddressRepository) Search(ctx ApiContext, criteria Criteria) (*CustomerAddressCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/customer-address", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(CustomerAddressCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t CustomerAddressRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/customer-address", criteria)

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

func (t CustomerAddressRepository) Upsert(ctx ApiContext, entity []CustomerAddress) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"customer_address": {
		Entity:  "customer_address",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t CustomerAddressRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"customer_address": {
		Entity:  "customer_address",
		Action:  "delete",
		Payload: payload,
	}})
}

type CustomerAddress struct {
	FirstName string `json:"firstName,omitempty"`

	Company string `json:"company,omitempty"`

	Department string `json:"department,omitempty"`

	AdditionalAddressLine1 string `json:"additionalAddressLine1,omitempty"`

	Id string `json:"id,omitempty"`

	CustomerId string `json:"customerId,omitempty"`

	CountryId string `json:"countryId,omitempty"`

	CountryStateId string `json:"countryStateId,omitempty"`

	Customer *Customer `json:"customer,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Salutation *Salutation `json:"salutation,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	LastName string `json:"lastName,omitempty"`

	City string `json:"city,omitempty"`

	Title string `json:"title,omitempty"`

	AdditionalAddressLine2 string `json:"additionalAddressLine2,omitempty"`

	Street string `json:"street,omitempty"`

	PhoneNumber string `json:"phoneNumber,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	SalutationId string `json:"salutationId,omitempty"`

	Zipcode string `json:"zipcode,omitempty"`

	Country *Country `json:"country,omitempty"`

	CountryState *CountryState `json:"countryState,omitempty"`
}

type CustomerAddressCollection struct {
	EntityCollection

	Data []CustomerAddress `json:"data"`
}

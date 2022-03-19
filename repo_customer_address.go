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

func (t CustomerAddressRepository) SearchAll(ctx ApiContext, criteria Criteria) (*CustomerAddressCollection, *http.Response, error) {
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
	SalutationId string `json:"salutationId,omitempty"`

	LastName string `json:"lastName,omitempty"`

	City string `json:"city,omitempty"`

	Department string `json:"department,omitempty"`

	AdditionalAddressLine1 string `json:"additionalAddressLine1,omitempty"`

	AdditionalAddressLine2 string `json:"additionalAddressLine2,omitempty"`

	Id string `json:"id,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	Customer *Customer `json:"customer,omitempty"`

	CountryState *CountryState `json:"countryState,omitempty"`

	Salutation *Salutation `json:"salutation,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Zipcode string `json:"zipcode,omitempty"`

	CountryStateId string `json:"countryStateId,omitempty"`

	Company string `json:"company,omitempty"`

	Street string `json:"street,omitempty"`

	PhoneNumber string `json:"phoneNumber,omitempty"`

	Country *Country `json:"country,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	CountryId string `json:"countryId,omitempty"`

	FirstName string `json:"firstName,omitempty"`

	Title string `json:"title,omitempty"`

	CustomerId string `json:"customerId,omitempty"`
}

type CustomerAddressCollection struct {
	EntityCollection

	Data []CustomerAddress `json:"data"`
}

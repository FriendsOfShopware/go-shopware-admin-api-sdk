package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type OrderAddressRepository ClientService

func (t OrderAddressRepository) Search(ctx ApiContext, criteria Criteria) (*OrderAddressCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/order-address", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(OrderAddressCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t OrderAddressRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/order-address", criteria)

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

func (t OrderAddressRepository) Upsert(ctx ApiContext, entity []OrderAddress) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"order_address": {
		Entity:  "order_address",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t OrderAddressRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"order_address": {
		Entity:  "order_address",
		Action:  "delete",
		Payload: payload,
	}})
}

type OrderAddress struct {
	CountryId string `json:"countryId,omitempty"`

	Title string `json:"title,omitempty"`

	AdditionalAddressLine1 string `json:"additionalAddressLine1,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	Country *Country `json:"country,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	LastName string `json:"lastName,omitempty"`

	CountryStateId string `json:"countryStateId,omitempty"`

	OrderId string `json:"orderId,omitempty"`

	City string `json:"city,omitempty"`

	PhoneNumber string `json:"phoneNumber,omitempty"`

	Id string `json:"id,omitempty"`

	Street string `json:"street,omitempty"`

	Order *Order `json:"order,omitempty"`

	Salutation *Salutation `json:"salutation,omitempty"`

	OrderVersionId string `json:"orderVersionId,omitempty"`

	Company string `json:"company,omitempty"`

	AdditionalAddressLine2 string `json:"additionalAddressLine2,omitempty"`

	CountryState *CountryState `json:"countryState,omitempty"`

	OrderDeliveries []OrderDelivery `json:"orderDeliveries,omitempty"`

	SalutationId string `json:"salutationId,omitempty"`

	VatId string `json:"vatId,omitempty"`

	VersionId string `json:"versionId,omitempty"`

	FirstName string `json:"firstName,omitempty"`

	Zipcode string `json:"zipcode,omitempty"`

	Department string `json:"department,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`
}

type OrderAddressCollection struct {
	EntityCollection

	Data []OrderAddress `json:"data"`
}

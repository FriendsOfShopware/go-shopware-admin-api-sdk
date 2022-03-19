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

func (t OrderAddressRepository) SearchAll(ctx ApiContext, criteria Criteria) (*OrderAddressCollection, *http.Response, error) {
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
	Street string `json:"street,omitempty"`

	VersionId string `json:"versionId,omitempty"`

	SalutationId string `json:"salutationId,omitempty"`

	FirstName string `json:"firstName,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Id string `json:"id,omitempty"`

	City string `json:"city,omitempty"`

	CountryState *CountryState `json:"countryState,omitempty"`

	Order *Order `json:"order,omitempty"`

	OrderVersionId string `json:"orderVersionId,omitempty"`

	LastName string `json:"lastName,omitempty"`

	Company string `json:"company,omitempty"`

	VatId string `json:"vatId,omitempty"`

	OrderDeliveries []OrderDelivery `json:"orderDeliveries,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	CountryStateId string `json:"countryStateId,omitempty"`

	Department string `json:"department,omitempty"`

	AdditionalAddressLine1 string `json:"additionalAddressLine1,omitempty"`

	Country *Country `json:"country,omitempty"`

	Zipcode string `json:"zipcode,omitempty"`

	PhoneNumber string `json:"phoneNumber,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	Salutation *Salutation `json:"salutation,omitempty"`

	CountryId string `json:"countryId,omitempty"`

	OrderId string `json:"orderId,omitempty"`

	Title string `json:"title,omitempty"`

	AdditionalAddressLine2 string `json:"additionalAddressLine2,omitempty"`
}

type OrderAddressCollection struct {
	EntityCollection

	Data []OrderAddress `json:"data"`
}

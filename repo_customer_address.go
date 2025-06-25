package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type CustomerAddressRepository struct {
	*GenericRepository[CustomerAddress]
}

func NewCustomerAddressRepository(client *Client) *CustomerAddressRepository {
	return &CustomerAddressRepository{
		GenericRepository: NewGenericRepository[CustomerAddress](client),
	}
}

func (t *CustomerAddressRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[CustomerAddress], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "customer-address")
}

func (t *CustomerAddressRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[CustomerAddress], *http.Response, error) {
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

func (t *CustomerAddressRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "customer-address")
}

func (t *CustomerAddressRepository) Upsert(ctx ApiContext, entity []CustomerAddress) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "customer_address")
}

func (t *CustomerAddressRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "customer_address")
}

type CustomerAddress struct {

	AdditionalAddressLine1      string  `json:"additionalAddressLine1,omitempty"`

	AdditionalAddressLine2      string  `json:"additionalAddressLine2,omitempty"`

	City      string  `json:"city,omitempty"`

	Company      string  `json:"company,omitempty"`

	Country      *Country  `json:"country,omitempty"`

	CountryId      string  `json:"countryId,omitempty"`

	CountryState      *CountryState  `json:"countryState,omitempty"`

	CountryStateId      string  `json:"countryStateId,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Customer      *Customer  `json:"customer,omitempty"`

	CustomerId      string  `json:"customerId,omitempty"`

	Department      string  `json:"department,omitempty"`

	FirstName      string  `json:"firstName,omitempty"`

	Hash      string  `json:"hash,omitempty"`

	Id      string  `json:"id,omitempty"`

	LastName      string  `json:"lastName,omitempty"`

	PhoneNumber      string  `json:"phoneNumber,omitempty"`

	Salutation      *Salutation  `json:"salutation,omitempty"`

	SalutationId      string  `json:"salutationId,omitempty"`

	Street      string  `json:"street,omitempty"`

	Title      string  `json:"title,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Zipcode      string  `json:"zipcode,omitempty"`

}

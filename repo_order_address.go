package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type OrderAddressRepository struct {
	*GenericRepository[OrderAddress]
}

func NewOrderAddressRepository(client *Client) *OrderAddressRepository {
	return &OrderAddressRepository{
		GenericRepository: NewGenericRepository[OrderAddress](client),
	}
}

func (t *OrderAddressRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[OrderAddress], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "order-address")
}

func (t *OrderAddressRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[OrderAddress], *http.Response, error) {
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

func (t *OrderAddressRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "order-address")
}

func (t *OrderAddressRepository) Upsert(ctx ApiContext, entity []OrderAddress) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "order_address")
}

func (t *OrderAddressRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "order_address")
}

type OrderAddress struct {

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

	Department      string  `json:"department,omitempty"`

	FirstName      string  `json:"firstName,omitempty"`

	Hash      string  `json:"hash,omitempty"`

	Id      string  `json:"id,omitempty"`

	LastName      string  `json:"lastName,omitempty"`

	Order      *Order  `json:"order,omitempty"`

	OrderDeliveries      []OrderDelivery  `json:"orderDeliveries,omitempty"`

	OrderId      string  `json:"orderId,omitempty"`

	OrderVersionId      string  `json:"orderVersionId,omitempty"`

	PhoneNumber      string  `json:"phoneNumber,omitempty"`

	Salutation      *Salutation  `json:"salutation,omitempty"`

	SalutationId      string  `json:"salutationId,omitempty"`

	Street      string  `json:"street,omitempty"`

	Title      string  `json:"title,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	VatId      string  `json:"vatId,omitempty"`

	VersionId      string  `json:"versionId,omitempty"`

	Zipcode      string  `json:"zipcode,omitempty"`

}

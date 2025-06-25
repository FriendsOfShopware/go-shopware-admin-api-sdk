package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type TaxRepository struct {
	*GenericRepository[Tax]
}

func NewTaxRepository(client *Client) *TaxRepository {
	return &TaxRepository{
		GenericRepository: NewGenericRepository[Tax](client),
	}
}

func (t *TaxRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[Tax], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "tax")
}

func (t *TaxRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[Tax], *http.Response, error) {
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

func (t *TaxRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "tax")
}

func (t *TaxRepository) Upsert(ctx ApiContext, entity []Tax) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "tax")
}

func (t *TaxRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "tax")
}

type Tax struct {

	Id      string  `json:"id,omitempty"`

	Name      string  `json:"name,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Products      []Product  `json:"products,omitempty"`

	ShippingMethods      []ShippingMethod  `json:"shippingMethods,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	TaxRate      float64  `json:"taxRate,omitempty"`

	Position      float64  `json:"position,omitempty"`

	Rules      []TaxRule  `json:"rules,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}

package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type DeliveryTimeRepository struct {
	*GenericRepository[DeliveryTime]
}

func NewDeliveryTimeRepository(client *Client) *DeliveryTimeRepository {
	return &DeliveryTimeRepository{
		GenericRepository: NewGenericRepository[DeliveryTime](client),
	}
}

func (t *DeliveryTimeRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[DeliveryTime], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "delivery-time")
}

func (t *DeliveryTimeRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[DeliveryTime], *http.Response, error) {
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

func (t *DeliveryTimeRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "delivery-time")
}

func (t *DeliveryTimeRepository) Upsert(ctx ApiContext, entity []DeliveryTime) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "delivery_time")
}

func (t *DeliveryTimeRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "delivery_time")
}

type DeliveryTime struct {

	Id      string  `json:"id,omitempty"`

	Name      string  `json:"name,omitempty"`

	Min      float64  `json:"min,omitempty"`

	Max      float64  `json:"max,omitempty"`

	Unit      string  `json:"unit,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Products      []Product  `json:"products,omitempty"`

	Translations      []DeliveryTimeTranslation  `json:"translations,omitempty"`

	ShippingMethods      []ShippingMethod  `json:"shippingMethods,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

}

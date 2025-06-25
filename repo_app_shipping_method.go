package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type AppShippingMethodRepository struct {
	*GenericRepository[AppShippingMethod]
}

func NewAppShippingMethodRepository(client *Client) *AppShippingMethodRepository {
	return &AppShippingMethodRepository{
		GenericRepository: NewGenericRepository[AppShippingMethod](client),
	}
}

func (t *AppShippingMethodRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[AppShippingMethod], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "app-shipping-method")
}

func (t *AppShippingMethodRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[AppShippingMethod], *http.Response, error) {
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

func (t *AppShippingMethodRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "app-shipping-method")
}

func (t *AppShippingMethodRepository) Upsert(ctx ApiContext, entity []AppShippingMethod) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "app_shipping_method")
}

func (t *AppShippingMethodRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "app_shipping_method")
}

type AppShippingMethod struct {

	App      *App  `json:"app,omitempty"`

	AppId      string  `json:"appId,omitempty"`

	AppName      string  `json:"appName,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Id      string  `json:"id,omitempty"`

	Identifier      string  `json:"identifier,omitempty"`

	OriginalMedia      *Media  `json:"originalMedia,omitempty"`

	OriginalMediaId      string  `json:"originalMediaId,omitempty"`

	ShippingMethod      *ShippingMethod  `json:"shippingMethod,omitempty"`

	ShippingMethodId      string  `json:"shippingMethodId,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}

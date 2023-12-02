package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type AppShippingMethodRepository ClientService

func (t AppShippingMethodRepository) Search(ctx ApiContext, criteria Criteria) (*AppShippingMethodCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/app-shipping-method", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(AppShippingMethodCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t AppShippingMethodRepository) SearchAll(ctx ApiContext, criteria Criteria) (*AppShippingMethodCollection, *http.Response, error) {
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

func (t AppShippingMethodRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/app-shipping-method", criteria)

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

func (t AppShippingMethodRepository) Upsert(ctx ApiContext, entity []AppShippingMethod) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"app_shipping_method": {
		Entity:  "app_shipping_method",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t AppShippingMethodRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"app_shipping_method": {
		Entity:  "app_shipping_method",
		Action:  "delete",
		Payload: payload,
	}})
}

type AppShippingMethod struct {

	AppId      string  `json:"appId,omitempty"`

	App      *App  `json:"app,omitempty"`

	ShippingMethodId      string  `json:"shippingMethodId,omitempty"`

	OriginalMediaId      string  `json:"originalMediaId,omitempty"`

	AppName      string  `json:"appName,omitempty"`

	Identifier      string  `json:"identifier,omitempty"`

	OriginalMedia      *Media  `json:"originalMedia,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Id      string  `json:"id,omitempty"`

	ShippingMethod      *ShippingMethod  `json:"shippingMethod,omitempty"`

}

type AppShippingMethodCollection struct {
	EntityCollection

	Data []AppShippingMethod `json:"data"`
}

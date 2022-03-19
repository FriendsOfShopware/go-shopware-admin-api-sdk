package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type AppPaymentMethodRepository ClientService

func (t AppPaymentMethodRepository) Search(ctx ApiContext, criteria Criteria) (*AppPaymentMethodCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/app-payment-method", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(AppPaymentMethodCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t AppPaymentMethodRepository) SearchAll(ctx ApiContext, criteria Criteria) (*AppPaymentMethodCollection, *http.Response, error) {
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

func (t AppPaymentMethodRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/app-payment-method", criteria)

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

func (t AppPaymentMethodRepository) Upsert(ctx ApiContext, entity []AppPaymentMethod) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"app_payment_method": {
		Entity:  "app_payment_method",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t AppPaymentMethodRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"app_payment_method": {
		Entity:  "app_payment_method",
		Action:  "delete",
		Payload: payload,
	}})
}

type AppPaymentMethod struct {
	AppName string `json:"appName,omitempty"`

	PayUrl string `json:"payUrl,omitempty"`

	AppId string `json:"appId,omitempty"`

	App *App `json:"app,omitempty"`

	OriginalMediaId string `json:"originalMediaId,omitempty"`

	PaymentMethod *PaymentMethod `json:"paymentMethod,omitempty"`

	Id string `json:"id,omitempty"`

	Identifier string `json:"identifier,omitempty"`

	FinalizeUrl string `json:"finalizeUrl,omitempty"`

	OriginalMedia *Media `json:"originalMedia,omitempty"`

	PaymentMethodId string `json:"paymentMethodId,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type AppPaymentMethodCollection struct {
	EntityCollection

	Data []AppPaymentMethod `json:"data"`
}

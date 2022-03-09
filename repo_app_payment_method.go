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
	PaymentMethod *PaymentMethod `json:"paymentMethod,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Id string `json:"id,omitempty"`

	PayUrl string `json:"payUrl,omitempty"`

	AppId string `json:"appId,omitempty"`

	App *App `json:"app,omitempty"`

	OriginalMediaId string `json:"originalMediaId,omitempty"`

	PaymentMethodId string `json:"paymentMethodId,omitempty"`

	AppName string `json:"appName,omitempty"`

	Identifier string `json:"identifier,omitempty"`

	FinalizeUrl string `json:"finalizeUrl,omitempty"`

	OriginalMedia *Media `json:"originalMedia,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`
}

type AppPaymentMethodCollection struct {
	EntityCollection

	Data []AppPaymentMethod `json:"data"`
}

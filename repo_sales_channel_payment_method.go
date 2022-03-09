package go_shopware_admin_sdk

import (
	"net/http"
)

type SalesChannelPaymentMethodRepository ClientService

func (t SalesChannelPaymentMethodRepository) Search(ctx ApiContext, criteria Criteria) (*SalesChannelPaymentMethodCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/sales-channel-payment-method", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(SalesChannelPaymentMethodCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t SalesChannelPaymentMethodRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/sales-channel-payment-method", criteria)

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

func (t SalesChannelPaymentMethodRepository) Upsert(ctx ApiContext, entity []SalesChannelPaymentMethod) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"sales_channel_payment_method": {
		Entity:  "sales_channel_payment_method",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t SalesChannelPaymentMethodRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"sales_channel_payment_method": {
		Entity:  "sales_channel_payment_method",
		Action:  "delete",
		Payload: payload,
	}})
}

type SalesChannelPaymentMethod struct {
	PaymentMethod *PaymentMethod `json:"paymentMethod,omitempty"`

	SalesChannelId string `json:"salesChannelId,omitempty"`

	PaymentMethodId string `json:"paymentMethodId,omitempty"`

	SalesChannel *SalesChannel `json:"salesChannel,omitempty"`
}

type SalesChannelPaymentMethodCollection struct {
	EntityCollection

	Data []SalesChannelPaymentMethod `json:"data"`
}

package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type PaymentMethodTranslationRepository ClientService

func (t PaymentMethodTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*PaymentMethodTranslationCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/payment-method-translation", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(PaymentMethodTranslationCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t PaymentMethodTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/payment-method-translation", criteria)

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

func (t PaymentMethodTranslationRepository) Upsert(ctx ApiContext, entity []PaymentMethodTranslation) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"payment_method_translation": {
		Entity:  "payment_method_translation",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t PaymentMethodTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"payment_method_translation": {
		Entity:  "payment_method_translation",
		Action:  "delete",
		Payload: payload,
	}})
}

type PaymentMethodTranslation struct {
	LanguageId string `json:"languageId,omitempty"`

	PaymentMethod *PaymentMethod `json:"paymentMethod,omitempty"`

	Language *Language `json:"language,omitempty"`

	Name string `json:"name,omitempty"`

	Description string `json:"description,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	PaymentMethodId string `json:"paymentMethodId,omitempty"`

	DistinguishableName string `json:"distinguishableName,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type PaymentMethodTranslationCollection struct {
	EntityCollection

	Data []PaymentMethodTranslation `json:"data"`
}

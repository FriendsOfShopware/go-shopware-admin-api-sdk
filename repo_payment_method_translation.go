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

func (t PaymentMethodTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*PaymentMethodTranslationCollection, *http.Response, error) {
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
	Name string `json:"name,omitempty"`

	PaymentMethodId string `json:"paymentMethodId,omitempty"`

	PaymentMethod *PaymentMethod `json:"paymentMethod,omitempty"`

	DistinguishableName string `json:"distinguishableName,omitempty"`

	Description string `json:"description,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	Language *Language `json:"language,omitempty"`
}

type PaymentMethodTranslationCollection struct {
	EntityCollection

	Data []PaymentMethodTranslation `json:"data"`
}

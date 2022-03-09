package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type ShippingMethodTranslationRepository ClientService

func (t ShippingMethodTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*ShippingMethodTranslationCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/shipping-method-translation", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(ShippingMethodTranslationCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ShippingMethodTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/shipping-method-translation", criteria)

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

func (t ShippingMethodTranslationRepository) Upsert(ctx ApiContext, entity []ShippingMethodTranslation) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"shipping_method_translation": {
		Entity:  "shipping_method_translation",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t ShippingMethodTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"shipping_method_translation": {
		Entity:  "shipping_method_translation",
		Action:  "delete",
		Payload: payload,
	}})
}

type ShippingMethodTranslation struct {
	LanguageId string `json:"languageId,omitempty"`

	ShippingMethod *ShippingMethod `json:"shippingMethod,omitempty"`

	Language *Language `json:"language,omitempty"`

	Name string `json:"name,omitempty"`

	Description string `json:"description,omitempty"`

	TrackingUrl string `json:"trackingUrl,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	ShippingMethodId string `json:"shippingMethodId,omitempty"`
}

type ShippingMethodTranslationCollection struct {
	EntityCollection

	Data []ShippingMethodTranslation `json:"data"`
}

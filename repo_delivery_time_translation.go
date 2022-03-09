package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type DeliveryTimeTranslationRepository ClientService

func (t DeliveryTimeTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*DeliveryTimeTranslationCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/delivery-time-translation", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(DeliveryTimeTranslationCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t DeliveryTimeTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/delivery-time-translation", criteria)

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

func (t DeliveryTimeTranslationRepository) Upsert(ctx ApiContext, entity []DeliveryTimeTranslation) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"delivery_time_translation": {
		Entity:  "delivery_time_translation",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t DeliveryTimeTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"delivery_time_translation": {
		Entity:  "delivery_time_translation",
		Action:  "delete",
		Payload: payload,
	}})
}

type DeliveryTimeTranslation struct {
	Language *Language `json:"language,omitempty"`

	Name string `json:"name,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	DeliveryTimeId string `json:"deliveryTimeId,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	DeliveryTime *DeliveryTime `json:"deliveryTime,omitempty"`
}

type DeliveryTimeTranslationCollection struct {
	EntityCollection

	Data []DeliveryTimeTranslation `json:"data"`
}

package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type ProductStreamTranslationRepository ClientService

func (t ProductStreamTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*ProductStreamTranslationCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/product-stream-translation", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(ProductStreamTranslationCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ProductStreamTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/product-stream-translation", criteria)

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

func (t ProductStreamTranslationRepository) Upsert(ctx ApiContext, entity []ProductStreamTranslation) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_stream_translation": {
		Entity:  "product_stream_translation",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t ProductStreamTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_stream_translation": {
		Entity:  "product_stream_translation",
		Action:  "delete",
		Payload: payload,
	}})
}

type ProductStreamTranslation struct {
	Language *Language `json:"language,omitempty"`

	Description string `json:"description,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	ProductStream *ProductStream `json:"productStream,omitempty"`

	Name string `json:"name,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	ProductStreamId string `json:"productStreamId,omitempty"`
}

type ProductStreamTranslationCollection struct {
	EntityCollection

	Data []ProductStreamTranslation `json:"data"`
}

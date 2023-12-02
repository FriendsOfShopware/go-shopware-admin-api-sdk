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

func (t ProductStreamTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*ProductStreamTranslationCollection, *http.Response, error) {
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

	ProductStream      *ProductStream  `json:"productStream,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	ProductStreamId      string  `json:"productStreamId,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	Name      string  `json:"name,omitempty"`

	Description      string  `json:"description,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

}

type ProductStreamTranslationCollection struct {
	EntityCollection

	Data []ProductStreamTranslation `json:"data"`
}

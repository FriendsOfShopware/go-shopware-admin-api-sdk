package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type NumberRangeTranslationRepository ClientService

func (t NumberRangeTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*NumberRangeTranslationCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/number-range-translation", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(NumberRangeTranslationCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t NumberRangeTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/number-range-translation", criteria)

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

func (t NumberRangeTranslationRepository) Upsert(ctx ApiContext, entity []NumberRangeTranslation) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"number_range_translation": {
		Entity:  "number_range_translation",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t NumberRangeTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"number_range_translation": {
		Entity:  "number_range_translation",
		Action:  "delete",
		Payload: payload,
	}})
}

type NumberRangeTranslation struct {
	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	Name string `json:"name,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	NumberRangeId string `json:"numberRangeId,omitempty"`

	NumberRange *NumberRange `json:"numberRange,omitempty"`

	Language *Language `json:"language,omitempty"`

	Description string `json:"description,omitempty"`
}

type NumberRangeTranslationCollection struct {
	EntityCollection

	Data []NumberRangeTranslation `json:"data"`
}

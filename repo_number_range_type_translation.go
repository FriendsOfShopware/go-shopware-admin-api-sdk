package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type NumberRangeTypeTranslationRepository ClientService

func (t NumberRangeTypeTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*NumberRangeTypeTranslationCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/number-range-type-translation", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(NumberRangeTypeTranslationCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t NumberRangeTypeTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*NumberRangeTypeTranslationCollection, *http.Response, error) {
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

func (t NumberRangeTypeTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/number-range-type-translation", criteria)

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

func (t NumberRangeTypeTranslationRepository) Upsert(ctx ApiContext, entity []NumberRangeTypeTranslation) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"number_range_type_translation": {
		Entity:  "number_range_type_translation",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t NumberRangeTypeTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"number_range_type_translation": {
		Entity:  "number_range_type_translation",
		Action:  "delete",
		Payload: payload,
	}})
}

type NumberRangeTypeTranslation struct {

	NumberRangeTypeId      string  `json:"numberRangeTypeId,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	NumberRangeType      *NumberRangeType  `json:"numberRangeType,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	TypeName      string  `json:"typeName,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}

type NumberRangeTypeTranslationCollection struct {
	EntityCollection

	Data []NumberRangeTypeTranslation `json:"data"`
}

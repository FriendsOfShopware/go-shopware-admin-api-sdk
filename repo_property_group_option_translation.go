package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type PropertyGroupOptionTranslationRepository ClientService

func (t PropertyGroupOptionTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*PropertyGroupOptionTranslationCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/property-group-option-translation", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(PropertyGroupOptionTranslationCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t PropertyGroupOptionTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*PropertyGroupOptionTranslationCollection, *http.Response, error) {
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

func (t PropertyGroupOptionTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/property-group-option-translation", criteria)

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

func (t PropertyGroupOptionTranslationRepository) Upsert(ctx ApiContext, entity []PropertyGroupOptionTranslation) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"property_group_option_translation": {
		Entity:  "property_group_option_translation",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t PropertyGroupOptionTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"property_group_option_translation": {
		Entity:  "property_group_option_translation",
		Action:  "delete",
		Payload: payload,
	}})
}

type PropertyGroupOptionTranslation struct {
	CustomFields interface{} `json:"customFields,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Language *Language `json:"language,omitempty"`

	Name string `json:"name,omitempty"`

	Position float64 `json:"position,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	PropertyGroupOptionId string `json:"propertyGroupOptionId,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	PropertyGroupOption *PropertyGroupOption `json:"propertyGroupOption,omitempty"`
}

type PropertyGroupOptionTranslationCollection struct {
	EntityCollection

	Data []PropertyGroupOptionTranslation `json:"data"`
}

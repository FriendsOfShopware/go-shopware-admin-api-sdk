package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type PropertyGroupTranslationRepository ClientService

func (t PropertyGroupTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*PropertyGroupTranslationCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/property-group-translation", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(PropertyGroupTranslationCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t PropertyGroupTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/property-group-translation", criteria)

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

func (t PropertyGroupTranslationRepository) Upsert(ctx ApiContext, entity []PropertyGroupTranslation) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"property_group_translation": {
		Entity:  "property_group_translation",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t PropertyGroupTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"property_group_translation": {
		Entity:  "property_group_translation",
		Action:  "delete",
		Payload: payload,
	}})
}

type PropertyGroupTranslation struct {
	Position float64 `json:"position,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	Language *Language `json:"language,omitempty"`

	PropertyGroupId string `json:"propertyGroupId,omitempty"`

	PropertyGroup *PropertyGroup `json:"propertyGroup,omitempty"`

	Name string `json:"name,omitempty"`

	Description string `json:"description,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type PropertyGroupTranslationCollection struct {
	EntityCollection

	Data []PropertyGroupTranslation `json:"data"`
}

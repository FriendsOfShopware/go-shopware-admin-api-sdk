package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type PropertyGroupOptionRepository ClientService

func (t PropertyGroupOptionRepository) Search(ctx ApiContext, criteria Criteria) (*PropertyGroupOptionCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/property-group-option", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(PropertyGroupOptionCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t PropertyGroupOptionRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/property-group-option", criteria)

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

func (t PropertyGroupOptionRepository) Upsert(ctx ApiContext, entity []PropertyGroupOption) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"property_group_option": {
		Entity:  "property_group_option",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t PropertyGroupOptionRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"property_group_option": {
		Entity:  "property_group_option",
		Action:  "delete",
		Payload: payload,
	}})
}

type PropertyGroupOption struct {
	Id string `json:"id,omitempty"`

	ProductProperties []Product `json:"productProperties,omitempty"`

	Name string `json:"name,omitempty"`

	Media *Media `json:"media,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Position float64 `json:"position,omitempty"`

	MediaId string `json:"mediaId,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	Translations []PropertyGroupOptionTranslation `json:"translations,omitempty"`

	ProductConfiguratorSettings []ProductConfiguratorSetting `json:"productConfiguratorSettings,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	GroupId string `json:"groupId,omitempty"`

	ColorHexCode string `json:"colorHexCode,omitempty"`

	Group *PropertyGroup `json:"group,omitempty"`

	ProductOptions []Product `json:"productOptions,omitempty"`

	Translated interface{} `json:"translated,omitempty"`
}

type PropertyGroupOptionCollection struct {
	EntityCollection

	Data []PropertyGroupOption `json:"data"`
}

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

func (t PropertyGroupOptionRepository) SearchAll(ctx ApiContext, criteria Criteria) (*PropertyGroupOptionCollection, *http.Response, error) {
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
	Group *PropertyGroup `json:"group,omitempty"`

	ProductConfiguratorSettings []ProductConfiguratorSetting `json:"productConfiguratorSettings,omitempty"`

	ProductProperties []Product `json:"productProperties,omitempty"`

	GroupId string `json:"groupId,omitempty"`

	Position float64 `json:"position,omitempty"`

	Media *Media `json:"media,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Name string `json:"name,omitempty"`

	MediaId string `json:"mediaId,omitempty"`

	Translations []PropertyGroupOptionTranslation `json:"translations,omitempty"`

	Id string `json:"id,omitempty"`

	ColorHexCode string `json:"colorHexCode,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	ProductOptions []Product `json:"productOptions,omitempty"`

	Translated interface{} `json:"translated,omitempty"`
}

type PropertyGroupOptionCollection struct {
	EntityCollection

	Data []PropertyGroupOption `json:"data"`
}

package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type PropertyGroupRepository ClientService

func (t PropertyGroupRepository) Search(ctx ApiContext, criteria Criteria) (*PropertyGroupCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/property-group", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(PropertyGroupCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t PropertyGroupRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/property-group", criteria)

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

func (t PropertyGroupRepository) Upsert(ctx ApiContext, entity []PropertyGroup) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"property_group": {
		Entity:  "property_group",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t PropertyGroupRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"property_group": {
		Entity:  "property_group",
		Action:  "delete",
		Payload: payload,
	}})
}

type PropertyGroup struct {
	Id string `json:"id,omitempty"`

	DisplayType string `json:"displayType,omitempty"`

	SortingType string `json:"sortingType,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Name string `json:"name,omitempty"`

	Description string `json:"description,omitempty"`

	Filterable bool `json:"filterable,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	VisibleOnProductDetailPage bool `json:"visibleOnProductDetailPage,omitempty"`

	Position float64 `json:"position,omitempty"`

	Translations []PropertyGroupTranslation `json:"translations,omitempty"`

	Options []PropertyGroupOption `json:"options,omitempty"`

	Translated interface{} `json:"translated,omitempty"`
}

type PropertyGroupCollection struct {
	EntityCollection

	Data []PropertyGroup `json:"data"`
}

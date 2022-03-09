package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type ProductConfiguratorSettingRepository ClientService

func (t ProductConfiguratorSettingRepository) Search(ctx ApiContext, criteria Criteria) (*ProductConfiguratorSettingCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/product-configurator-setting", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(ProductConfiguratorSettingCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ProductConfiguratorSettingRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/product-configurator-setting", criteria)

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

func (t ProductConfiguratorSettingRepository) Upsert(ctx ApiContext, entity []ProductConfiguratorSetting) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_configurator_setting": {
		Entity:  "product_configurator_setting",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t ProductConfiguratorSettingRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_configurator_setting": {
		Entity:  "product_configurator_setting",
		Action:  "delete",
		Payload: payload,
	}})
}

type ProductConfiguratorSetting struct {
	OptionId string `json:"optionId,omitempty"`

	Media *Media `json:"media,omitempty"`

	Option *PropertyGroupOption `json:"option,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Id string `json:"id,omitempty"`

	ProductVersionId string `json:"productVersionId,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	VersionId string `json:"versionId,omitempty"`

	Position float64 `json:"position,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	ProductId string `json:"productId,omitempty"`

	MediaId string `json:"mediaId,omitempty"`

	Price interface{} `json:"price,omitempty"`

	Product *Product `json:"product,omitempty"`
}

type ProductConfiguratorSettingCollection struct {
	EntityCollection

	Data []ProductConfiguratorSetting `json:"data"`
}

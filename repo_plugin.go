package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type PluginRepository ClientService

func (t PluginRepository) Search(ctx ApiContext, criteria Criteria) (*PluginCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/plugin", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(PluginCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t PluginRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/plugin", criteria)

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

func (t PluginRepository) Upsert(ctx ApiContext, entity []Plugin) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"plugin": {
		Entity:  "plugin",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t PluginRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"plugin": {
		Entity:  "plugin",
		Action:  "delete",
		Payload: payload,
	}})
}

type Plugin struct {
	ManagedByComposer bool `json:"managedByComposer,omitempty"`

	Changelog interface{} `json:"changelog,omitempty"`

	Translated interface{} `json:"translated,omitempty"`

	Version string `json:"version,omitempty"`

	UpgradedAt time.Time `json:"upgradedAt,omitempty"`

	IconRaw interface{} `json:"iconRaw,omitempty"`

	Icon string `json:"icon,omitempty"`

	Label string `json:"label,omitempty"`

	BaseClass string `json:"baseClass,omitempty"`

	ComposerName string `json:"composerName,omitempty"`

	Active bool `json:"active,omitempty"`

	ManufacturerLink string `json:"manufacturerLink,omitempty"`

	Translations []PluginTranslation `json:"translations,omitempty"`

	Author string `json:"author,omitempty"`

	SupportLink string `json:"supportLink,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	Description string `json:"description,omitempty"`

	PaymentMethods []PaymentMethod `json:"paymentMethods,omitempty"`

	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Autoload interface{} `json:"autoload,omitempty"`

	Copyright string `json:"copyright,omitempty"`

	UpgradeVersion string `json:"upgradeVersion,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Path string `json:"path,omitempty"`

	License string `json:"license,omitempty"`

	InstalledAt time.Time `json:"installedAt,omitempty"`
}

type PluginCollection struct {
	EntityCollection

	Data []Plugin `json:"data"`
}

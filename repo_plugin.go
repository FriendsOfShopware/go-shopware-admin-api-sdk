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

func (t PluginRepository) SearchAll(ctx ApiContext, criteria Criteria) (*PluginCollection, *http.Response, error) {
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

	Copyright      string  `json:"copyright,omitempty"`

	Version      string  `json:"version,omitempty"`

	Translations      []PluginTranslation  `json:"translations,omitempty"`

	ComposerName      string  `json:"composerName,omitempty"`

	Author      string  `json:"author,omitempty"`

	License      string  `json:"license,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	ManagedByComposer      bool  `json:"managedByComposer,omitempty"`

	Path      string  `json:"path,omitempty"`

	SupportLink      string  `json:"supportLink,omitempty"`

	Changelog      interface{}  `json:"changelog,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Name      string  `json:"name,omitempty"`

	Active      bool  `json:"active,omitempty"`

	UpgradeVersion      string  `json:"upgradeVersion,omitempty"`

	UpgradedAt      time.Time  `json:"upgradedAt,omitempty"`

	Icon      string  `json:"icon,omitempty"`

	Description      string  `json:"description,omitempty"`

	BaseClass      string  `json:"baseClass,omitempty"`

	InstalledAt      time.Time  `json:"installedAt,omitempty"`

	IconRaw      interface{}  `json:"iconRaw,omitempty"`

	Label      string  `json:"label,omitempty"`

	PaymentMethods      []PaymentMethod  `json:"paymentMethods,omitempty"`

	Id      string  `json:"id,omitempty"`

	Autoload      interface{}  `json:"autoload,omitempty"`

	ManufacturerLink      string  `json:"manufacturerLink,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

}

type PluginCollection struct {
	EntityCollection

	Data []Plugin `json:"data"`
}

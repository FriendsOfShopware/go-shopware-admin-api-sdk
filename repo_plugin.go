package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type PluginRepository struct {
	*GenericRepository[Plugin]
}

func NewPluginRepository(client *Client) *PluginRepository {
	return &PluginRepository{
		GenericRepository: NewGenericRepository[Plugin](client),
	}
}

func (t *PluginRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[Plugin], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "plugin")
}

func (t *PluginRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[Plugin], *http.Response, error) {
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

func (t *PluginRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "plugin")
}

func (t *PluginRepository) Upsert(ctx ApiContext, entity []Plugin) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "plugin")
}

func (t *PluginRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "plugin")
}

type Plugin struct {

	Label      string  `json:"label,omitempty"`

	ManufacturerLink      string  `json:"manufacturerLink,omitempty"`

	Translations      []PluginTranslation  `json:"translations,omitempty"`

	License      string  `json:"license,omitempty"`

	InstalledAt      time.Time  `json:"installedAt,omitempty"`

	SupportLink      string  `json:"supportLink,omitempty"`

	PaymentMethods      []PaymentMethod  `json:"paymentMethods,omitempty"`

	Id      string  `json:"id,omitempty"`

	BaseClass      string  `json:"baseClass,omitempty"`

	ComposerName      string  `json:"composerName,omitempty"`

	Copyright      string  `json:"copyright,omitempty"`

	Version      string  `json:"version,omitempty"`

	IconRaw      interface{}  `json:"iconRaw,omitempty"`

	Description      string  `json:"description,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Name      string  `json:"name,omitempty"`

	Active      bool  `json:"active,omitempty"`

	ManagedByComposer      bool  `json:"managedByComposer,omitempty"`

	UpgradeVersion      string  `json:"upgradeVersion,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	Autoload      interface{}  `json:"autoload,omitempty"`

	Path      string  `json:"path,omitempty"`

	Author      string  `json:"author,omitempty"`

	UpgradedAt      time.Time  `json:"upgradedAt,omitempty"`

	Icon      string  `json:"icon,omitempty"`

}

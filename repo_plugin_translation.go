package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type PluginTranslationRepository struct {
	*GenericRepository[PluginTranslation]
}

func NewPluginTranslationRepository(client *Client) *PluginTranslationRepository {
	return &PluginTranslationRepository{
		GenericRepository: NewGenericRepository[PluginTranslation](client),
	}
}

func (t *PluginTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[PluginTranslation], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "plugin-translation")
}

func (t *PluginTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[PluginTranslation], *http.Response, error) {
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

func (t *PluginTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "plugin-translation")
}

func (t *PluginTranslationRepository) Upsert(ctx ApiContext, entity []PluginTranslation) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "plugin_translation")
}

func (t *PluginTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "plugin_translation")
}

type PluginTranslation struct {

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Description      string  `json:"description,omitempty"`

	Label      string  `json:"label,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	ManufacturerLink      string  `json:"manufacturerLink,omitempty"`

	Plugin      *Plugin  `json:"plugin,omitempty"`

	PluginId      string  `json:"pluginId,omitempty"`

	SupportLink      string  `json:"supportLink,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}

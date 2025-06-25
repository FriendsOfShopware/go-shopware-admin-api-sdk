package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type PropertyGroupOptionRepository struct {
	*GenericRepository[PropertyGroupOption]
}

func NewPropertyGroupOptionRepository(client *Client) *PropertyGroupOptionRepository {
	return &PropertyGroupOptionRepository{
		GenericRepository: NewGenericRepository[PropertyGroupOption](client),
	}
}

func (t *PropertyGroupOptionRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[PropertyGroupOption], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "property-group-option")
}

func (t *PropertyGroupOptionRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[PropertyGroupOption], *http.Response, error) {
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

func (t *PropertyGroupOptionRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "property-group-option")
}

func (t *PropertyGroupOptionRepository) Upsert(ctx ApiContext, entity []PropertyGroupOption) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "property_group_option")
}

func (t *PropertyGroupOptionRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "property_group_option")
}

type PropertyGroupOption struct {

	Position      float64  `json:"position,omitempty"`

	Media      *Media  `json:"media,omitempty"`

	Translations      []PropertyGroupOptionTranslation  `json:"translations,omitempty"`

	Id      string  `json:"id,omitempty"`

	GroupId      string  `json:"groupId,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Group      *PropertyGroup  `json:"group,omitempty"`

	ProductOptions      []Product  `json:"productOptions,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	Name      string  `json:"name,omitempty"`

	ColorHexCode      string  `json:"colorHexCode,omitempty"`

	MediaId      string  `json:"mediaId,omitempty"`

	ProductConfiguratorSettings      []ProductConfiguratorSetting  `json:"productConfiguratorSettings,omitempty"`

	ProductProperties      []Product  `json:"productProperties,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}

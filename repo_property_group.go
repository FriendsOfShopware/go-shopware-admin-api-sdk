package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type PropertyGroupRepository struct {
	*GenericRepository[PropertyGroup]
}

func NewPropertyGroupRepository(client *Client) *PropertyGroupRepository {
	return &PropertyGroupRepository{
		GenericRepository: NewGenericRepository[PropertyGroup](client),
	}
}

func (t *PropertyGroupRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[PropertyGroup], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "property-group")
}

func (t *PropertyGroupRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[PropertyGroup], *http.Response, error) {
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

func (t *PropertyGroupRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "property-group")
}

func (t *PropertyGroupRepository) Upsert(ctx ApiContext, entity []PropertyGroup) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "property_group")
}

func (t *PropertyGroupRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "property_group")
}

type PropertyGroup struct {

	Translated      interface{}  `json:"translated,omitempty"`

	Id      string  `json:"id,omitempty"`

	Description      string  `json:"description,omitempty"`

	DisplayType      string  `json:"displayType,omitempty"`

	SortingType      string  `json:"sortingType,omitempty"`

	Filterable      bool  `json:"filterable,omitempty"`

	VisibleOnProductDetailPage      bool  `json:"visibleOnProductDetailPage,omitempty"`

	Position      float64  `json:"position,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Name      string  `json:"name,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Options      []PropertyGroupOption  `json:"options,omitempty"`

	Translations      []PropertyGroupTranslation  `json:"translations,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

}

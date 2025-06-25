package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type CustomEntityRepository struct {
	*GenericRepository[CustomEntity]
}

func NewCustomEntityRepository(client *Client) *CustomEntityRepository {
	return &CustomEntityRepository{
		GenericRepository: NewGenericRepository[CustomEntity](client),
	}
}

func (t *CustomEntityRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[CustomEntity], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "custom-entity")
}

func (t *CustomEntityRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[CustomEntity], *http.Response, error) {
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

func (t *CustomEntityRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "custom-entity")
}

func (t *CustomEntityRepository) Upsert(ctx ApiContext, entity []CustomEntity) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "custom_entity")
}

func (t *CustomEntityRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "custom_entity")
}

type CustomEntity struct {

	AppId      string  `json:"appId,omitempty"`

	CmsAware      bool  `json:"cmsAware,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFieldsAware      bool  `json:"customFieldsAware,omitempty"`

	DeletedAt      time.Time  `json:"deletedAt,omitempty"`

	Fields      interface{}  `json:"fields,omitempty"`

	Flags      interface{}  `json:"flags,omitempty"`

	Id      string  `json:"id,omitempty"`

	LabelProperty      string  `json:"labelProperty,omitempty"`

	Name      string  `json:"name,omitempty"`

	PluginId      string  `json:"pluginId,omitempty"`

	StoreApiAware      bool  `json:"storeApiAware,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}

package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type CustomFieldSetRepository struct {
	*GenericRepository[CustomFieldSet]
}

func NewCustomFieldSetRepository(client *Client) *CustomFieldSetRepository {
	return &CustomFieldSetRepository{
		GenericRepository: NewGenericRepository[CustomFieldSet](client),
	}
}

func (t *CustomFieldSetRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[CustomFieldSet], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "custom-field-set")
}

func (t *CustomFieldSetRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[CustomFieldSet], *http.Response, error) {
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

func (t *CustomFieldSetRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "custom-field-set")
}

func (t *CustomFieldSetRepository) Upsert(ctx ApiContext, entity []CustomFieldSet) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "custom_field_set")
}

func (t *CustomFieldSetRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "custom_field_set")
}

type CustomFieldSet struct {

	Active      bool  `json:"active,omitempty"`

	App      *App  `json:"app,omitempty"`

	AppId      string  `json:"appId,omitempty"`

	Config      interface{}  `json:"config,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFields      []CustomField  `json:"customFields,omitempty"`

	Global      bool  `json:"global,omitempty"`

	Id      string  `json:"id,omitempty"`

	Name      string  `json:"name,omitempty"`

	Position      float64  `json:"position,omitempty"`

	Products      []Product  `json:"products,omitempty"`

	Relations      []CustomFieldSetRelation  `json:"relations,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}

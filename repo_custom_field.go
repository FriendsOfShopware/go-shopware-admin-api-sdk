package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type CustomFieldRepository struct {
	*GenericRepository[CustomField]
}

func NewCustomFieldRepository(client *Client) *CustomFieldRepository {
	return &CustomFieldRepository{
		GenericRepository: NewGenericRepository[CustomField](client),
	}
}

func (t *CustomFieldRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[CustomField], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "custom-field")
}

func (t *CustomFieldRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[CustomField], *http.Response, error) {
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

func (t *CustomFieldRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "custom-field")
}

func (t *CustomFieldRepository) Upsert(ctx ApiContext, entity []CustomField) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "custom_field")
}

func (t *CustomFieldRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "custom_field")
}

type CustomField struct {

	Active      bool  `json:"active,omitempty"`

	AllowCartExpose      bool  `json:"allowCartExpose,omitempty"`

	AllowCustomerWrite      bool  `json:"allowCustomerWrite,omitempty"`

	Config      interface{}  `json:"config,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFieldSet      *CustomFieldSet  `json:"customFieldSet,omitempty"`

	CustomFieldSetId      string  `json:"customFieldSetId,omitempty"`

	Id      string  `json:"id,omitempty"`

	Name      string  `json:"name,omitempty"`

	ProductSearchConfigFields      []ProductSearchConfigField  `json:"productSearchConfigFields,omitempty"`

	StoreApiAware      bool  `json:"storeApiAware,omitempty"`

	Type      string  `json:"type,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}

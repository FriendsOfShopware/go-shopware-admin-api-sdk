package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type CustomFieldSetRelationRepository struct {
	*GenericRepository[CustomFieldSetRelation]
}

func NewCustomFieldSetRelationRepository(client *Client) *CustomFieldSetRelationRepository {
	return &CustomFieldSetRelationRepository{
		GenericRepository: NewGenericRepository[CustomFieldSetRelation](client),
	}
}

func (t *CustomFieldSetRelationRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[CustomFieldSetRelation], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "custom-field-set-relation")
}

func (t *CustomFieldSetRelationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[CustomFieldSetRelation], *http.Response, error) {
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

func (t *CustomFieldSetRelationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "custom-field-set-relation")
}

func (t *CustomFieldSetRelationRepository) Upsert(ctx ApiContext, entity []CustomFieldSetRelation) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "custom_field_set_relation")
}

func (t *CustomFieldSetRelationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "custom_field_set_relation")
}

type CustomFieldSetRelation struct {

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFieldSet      *CustomFieldSet  `json:"customFieldSet,omitempty"`

	CustomFieldSetId      string  `json:"customFieldSetId,omitempty"`

	EntityName      string  `json:"entityName,omitempty"`

	Id      string  `json:"id,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}

package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type CustomFieldSetRelationRepository ClientService

func (t CustomFieldSetRelationRepository) Search(ctx ApiContext, criteria Criteria) (*CustomFieldSetRelationCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/custom-field-set-relation", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(CustomFieldSetRelationCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t CustomFieldSetRelationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*CustomFieldSetRelationCollection, *http.Response, error) {
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

func (t CustomFieldSetRelationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/custom-field-set-relation", criteria)

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

func (t CustomFieldSetRelationRepository) Upsert(ctx ApiContext, entity []CustomFieldSetRelation) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"custom_field_set_relation": {
		Entity:  "custom_field_set_relation",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t CustomFieldSetRelationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"custom_field_set_relation": {
		Entity:  "custom_field_set_relation",
		Action:  "delete",
		Payload: payload,
	}})
}

type CustomFieldSetRelation struct {

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Id      string  `json:"id,omitempty"`

	CustomFieldSetId      string  `json:"customFieldSetId,omitempty"`

	EntityName      string  `json:"entityName,omitempty"`

	CustomFieldSet      *CustomFieldSet  `json:"customFieldSet,omitempty"`

}

type CustomFieldSetRelationCollection struct {
	EntityCollection

	Data []CustomFieldSetRelation `json:"data"`
}

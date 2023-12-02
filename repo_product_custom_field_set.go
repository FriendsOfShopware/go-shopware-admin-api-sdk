package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type ProductCustomFieldSetRepository ClientService

func (t ProductCustomFieldSetRepository) Search(ctx ApiContext, criteria Criteria) (*ProductCustomFieldSetCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/product-custom-field-set", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(ProductCustomFieldSetCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ProductCustomFieldSetRepository) SearchAll(ctx ApiContext, criteria Criteria) (*ProductCustomFieldSetCollection, *http.Response, error) {
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

func (t ProductCustomFieldSetRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/product-custom-field-set", criteria)

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

func (t ProductCustomFieldSetRepository) Upsert(ctx ApiContext, entity []ProductCustomFieldSet) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_custom_field_set": {
		Entity:  "product_custom_field_set",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t ProductCustomFieldSetRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_custom_field_set": {
		Entity:  "product_custom_field_set",
		Action:  "delete",
		Payload: payload,
	}})
}

type ProductCustomFieldSet struct {

	ProductVersionId      string  `json:"productVersionId,omitempty"`

	Product      *Product  `json:"product,omitempty"`

	CustomFieldSet      *CustomFieldSet  `json:"customFieldSet,omitempty"`

	ProductId      string  `json:"productId,omitempty"`

	CustomFieldSetId      string  `json:"customFieldSetId,omitempty"`

}

type ProductCustomFieldSetCollection struct {
	EntityCollection

	Data []ProductCustomFieldSet `json:"data"`
}

package go_shopware_admin_sdk

import (
	"net/http"
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
	Product *Product `json:"product,omitempty"`

	CustomFieldSet *CustomFieldSet `json:"customFieldSet,omitempty"`

	ProductId string `json:"productId,omitempty"`

	CustomFieldSetId string `json:"customFieldSetId,omitempty"`

	ProductVersionId string `json:"productVersionId,omitempty"`
}

type ProductCustomFieldSetCollection struct {
	EntityCollection

	Data []ProductCustomFieldSet `json:"data"`
}

package go_shopware_admin_sdk

import (
	"net/http"
)

type ShippingMethodTagRepository ClientService

func (t ShippingMethodTagRepository) Search(ctx ApiContext, criteria Criteria) (*ShippingMethodTagCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/shipping-method-tag", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(ShippingMethodTagCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ShippingMethodTagRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/shipping-method-tag", criteria)

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

func (t ShippingMethodTagRepository) Upsert(ctx ApiContext, entity []ShippingMethodTag) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"shipping_method_tag": {
		Entity:  "shipping_method_tag",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t ShippingMethodTagRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"shipping_method_tag": {
		Entity:  "shipping_method_tag",
		Action:  "delete",
		Payload: payload,
	}})
}

type ShippingMethodTag struct {
	ShippingMethodId string `json:"shippingMethodId,omitempty"`

	TagId string `json:"tagId,omitempty"`

	ShippingMethod *ShippingMethod `json:"shippingMethod,omitempty"`

	Tag *Tag `json:"tag,omitempty"`
}

type ShippingMethodTagCollection struct {
	EntityCollection

	Data []ShippingMethodTag `json:"data"`
}

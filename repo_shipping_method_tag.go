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

func (t ShippingMethodTagRepository) SearchAll(ctx ApiContext, criteria Criteria) (*ShippingMethodTagCollection, *http.Response, error) {
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

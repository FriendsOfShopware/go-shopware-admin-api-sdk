package go_shopware_admin_sdk

import (
	"net/http"
)

type OrderTagRepository ClientService

func (t OrderTagRepository) Search(ctx ApiContext, criteria Criteria) (*OrderTagCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/order-tag", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(OrderTagCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t OrderTagRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/order-tag", criteria)

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

func (t OrderTagRepository) Upsert(ctx ApiContext, entity []OrderTag) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"order_tag": {
		Entity:  "order_tag",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t OrderTagRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"order_tag": {
		Entity:  "order_tag",
		Action:  "delete",
		Payload: payload,
	}})
}

type OrderTag struct {
	Tag *Tag `json:"tag,omitempty"`

	OrderId string `json:"orderId,omitempty"`

	OrderVersionId string `json:"orderVersionId,omitempty"`

	TagId string `json:"tagId,omitempty"`

	Order *Order `json:"order,omitempty"`
}

type OrderTagCollection struct {
	EntityCollection

	Data []OrderTag `json:"data"`
}

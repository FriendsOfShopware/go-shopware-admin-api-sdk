package go_shopware_admin_sdk

import (
	"net/http"

	"time"
)

type OrderLineItemDownloadRepository ClientService

func (t OrderLineItemDownloadRepository) Search(ctx ApiContext, criteria Criteria) (*OrderLineItemDownloadCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/order-line-item-download", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(OrderLineItemDownloadCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t OrderLineItemDownloadRepository) SearchAll(ctx ApiContext, criteria Criteria) (*OrderLineItemDownloadCollection, *http.Response, error) {
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

func (t OrderLineItemDownloadRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/order-line-item-download", criteria)

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

func (t OrderLineItemDownloadRepository) Upsert(ctx ApiContext, entity []OrderLineItemDownload) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"order_line_item_download": {
		Entity:  "order_line_item_download",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t OrderLineItemDownloadRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"order_line_item_download": {
		Entity:  "order_line_item_download",
		Action:  "delete",
		Payload: payload,
	}})
}

type OrderLineItemDownload struct {
	Id string `json:"id,omitempty"`

	VersionId string `json:"versionId,omitempty"`

	MediaId string `json:"mediaId,omitempty"`

	OrderLineItem *OrderLineItem `json:"orderLineItem,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	OrderLineItemId string `json:"orderLineItemId,omitempty"`

	OrderLineItemVersionId string `json:"orderLineItemVersionId,omitempty"`

	Position float64 `json:"position,omitempty"`

	AccessGranted bool `json:"accessGranted,omitempty"`

	Media *Media `json:"media,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`
}

type OrderLineItemDownloadCollection struct {
	EntityCollection

	Data []OrderLineItemDownload `json:"data"`
}

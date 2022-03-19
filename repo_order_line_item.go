package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type OrderLineItemRepository ClientService

func (t OrderLineItemRepository) Search(ctx ApiContext, criteria Criteria) (*OrderLineItemCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/order-line-item", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(OrderLineItemCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t OrderLineItemRepository) SearchAll(ctx ApiContext, criteria Criteria) (*OrderLineItemCollection, *http.Response, error) {
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

func (t OrderLineItemRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/order-line-item", criteria)

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

func (t OrderLineItemRepository) Upsert(ctx ApiContext, entity []OrderLineItem) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"order_line_item": {
		Entity:  "order_line_item",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t OrderLineItemRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"order_line_item": {
		Entity:  "order_line_item",
		Action:  "delete",
		Payload: payload,
	}})
}

type OrderLineItem struct {
	ProductId string `json:"productId,omitempty"`

	ProductVersionId string `json:"productVersionId,omitempty"`

	Product *Product `json:"product,omitempty"`

	CoverId string `json:"coverId,omitempty"`

	ReferencedId string `json:"referencedId,omitempty"`

	Order *Order `json:"order,omitempty"`

	OrderVersionId string `json:"orderVersionId,omitempty"`

	Quantity float64 `json:"quantity,omitempty"`

	Position float64 `json:"position,omitempty"`

	Description string `json:"description,omitempty"`

	Price interface{} `json:"price,omitempty"`

	Children []OrderLineItem `json:"children,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	OrderId string `json:"orderId,omitempty"`

	PriceDefinition interface{} `json:"priceDefinition,omitempty"`

	UnitPrice float64 `json:"unitPrice,omitempty"`

	OrderDeliveryPositions []OrderDeliveryPosition `json:"orderDeliveryPositions,omitempty"`

	ParentId string `json:"parentId,omitempty"`

	Cover *Media `json:"cover,omitempty"`

	Label string `json:"label,omitempty"`

	Type string `json:"type,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	Parent *OrderLineItem `json:"parent,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Stackable bool `json:"stackable,omitempty"`

	Id string `json:"id,omitempty"`

	VersionId string `json:"versionId,omitempty"`

	ParentVersionId string `json:"parentVersionId,omitempty"`

	Identifier string `json:"identifier,omitempty"`

	Payload interface{} `json:"payload,omitempty"`

	Good bool `json:"good,omitempty"`

	Removable bool `json:"removable,omitempty"`

	TotalPrice float64 `json:"totalPrice,omitempty"`
}

type OrderLineItemCollection struct {
	EntityCollection

	Data []OrderLineItem `json:"data"`
}

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
	OrderVersionId string `json:"orderVersionId,omitempty"`

	ParentId string `json:"parentId,omitempty"`

	CoverId string `json:"coverId,omitempty"`

	Good bool `json:"good,omitempty"`

	OrderDeliveryPositions []OrderDeliveryPosition `json:"orderDeliveryPositions,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	OrderId string `json:"orderId,omitempty"`

	Quantity float64 `json:"quantity,omitempty"`

	TotalPrice float64 `json:"totalPrice,omitempty"`

	Type string `json:"type,omitempty"`

	Position float64 `json:"position,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Id string `json:"id,omitempty"`

	ProductVersionId string `json:"productVersionId,omitempty"`

	Identifier string `json:"identifier,omitempty"`

	Stackable bool `json:"stackable,omitempty"`

	ParentVersionId string `json:"parentVersionId,omitempty"`

	Label string `json:"label,omitempty"`

	Price interface{} `json:"price,omitempty"`

	Product *Product `json:"product,omitempty"`

	VersionId string `json:"versionId,omitempty"`

	ProductId string `json:"productId,omitempty"`

	Cover *Media `json:"cover,omitempty"`

	Removable bool `json:"removable,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	PriceDefinition interface{} `json:"priceDefinition,omitempty"`

	UnitPrice float64 `json:"unitPrice,omitempty"`

	Description string `json:"description,omitempty"`

	ReferencedId string `json:"referencedId,omitempty"`

	Payload interface{} `json:"payload,omitempty"`

	Order *Order `json:"order,omitempty"`

	Parent *OrderLineItem `json:"parent,omitempty"`

	Children []OrderLineItem `json:"children,omitempty"`
}

type OrderLineItemCollection struct {
	EntityCollection

	Data []OrderLineItem `json:"data"`
}

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
	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	ProductVersionId string `json:"productVersionId,omitempty"`

	UnitPrice float64 `json:"unitPrice,omitempty"`

	OrderDeliveryPositions []OrderDeliveryPosition `json:"orderDeliveryPositions,omitempty"`

	PriceDefinition interface{} `json:"priceDefinition,omitempty"`

	OrderVersionId string `json:"orderVersionId,omitempty"`

	Payload interface{} `json:"payload,omitempty"`

	Good bool `json:"good,omitempty"`

	ParentVersionId string `json:"parentVersionId,omitempty"`

	Position float64 `json:"position,omitempty"`

	Type string `json:"type,omitempty"`

	Product *Product `json:"product,omitempty"`

	Id string `json:"id,omitempty"`

	VersionId string `json:"versionId,omitempty"`

	PromotionId string `json:"promotionId,omitempty"`

	Stackable bool `json:"stackable,omitempty"`

	States interface{} `json:"states,omitempty"`

	Description string `json:"description,omitempty"`

	OrderId string `json:"orderId,omitempty"`

	ParentId string `json:"parentId,omitempty"`

	CoverId string `json:"coverId,omitempty"`

	OrderTransactionCaptureRefundPositions []OrderTransactionCaptureRefundPosition `json:"orderTransactionCaptureRefundPositions,omitempty"`

	Children []OrderLineItem `json:"children,omitempty"`

	Cover *Media `json:"cover,omitempty"`

	Identifier string `json:"identifier,omitempty"`

	ReferencedId string `json:"referencedId,omitempty"`

	Promotion *Promotion `json:"promotion,omitempty"`

	Downloads []OrderLineItemDownload `json:"downloads,omitempty"`

	Parent *OrderLineItem `json:"parent,omitempty"`

	Quantity float64 `json:"quantity,omitempty"`

	Removable bool `json:"removable,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	Price interface{} `json:"price,omitempty"`

	Order *Order `json:"order,omitempty"`

	ProductId string `json:"productId,omitempty"`

	Label string `json:"label,omitempty"`

	TotalPrice float64 `json:"totalPrice,omitempty"`
}

type OrderLineItemCollection struct {
	EntityCollection

	Data []OrderLineItem `json:"data"`
}

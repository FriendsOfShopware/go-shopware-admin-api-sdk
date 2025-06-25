package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type OrderLineItemRepository struct {
	*GenericRepository[OrderLineItem]
}

func NewOrderLineItemRepository(client *Client) *OrderLineItemRepository {
	return &OrderLineItemRepository{
		GenericRepository: NewGenericRepository[OrderLineItem](client),
	}
}

func (t *OrderLineItemRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[OrderLineItem], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "order-line-item")
}

func (t *OrderLineItemRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[OrderLineItem], *http.Response, error) {
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

func (t *OrderLineItemRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "order-line-item")
}

func (t *OrderLineItemRepository) Upsert(ctx ApiContext, entity []OrderLineItem) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "order_line_item")
}

func (t *OrderLineItemRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "order_line_item")
}

type OrderLineItem struct {

	ReferencedId      string  `json:"referencedId,omitempty"`

	States      interface{}  `json:"states,omitempty"`

	TotalPrice      float64  `json:"totalPrice,omitempty"`

	Type      string  `json:"type,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	OrderId      string  `json:"orderId,omitempty"`

	ProductId      string  `json:"productId,omitempty"`

	Cover      *Media  `json:"cover,omitempty"`

	PriceDefinition      interface{}  `json:"priceDefinition,omitempty"`

	CoverId      string  `json:"coverId,omitempty"`

	Stackable      bool  `json:"stackable,omitempty"`

	Price      interface{}  `json:"price,omitempty"`

	Order      *Order  `json:"order,omitempty"`

	Product      *Product  `json:"product,omitempty"`

	OrderDeliveryPositions      []OrderDeliveryPosition  `json:"orderDeliveryPositions,omitempty"`

	OrderTransactionCaptureRefundPositions      []OrderTransactionCaptureRefundPosition  `json:"orderTransactionCaptureRefundPositions,omitempty"`

	Identifier      string  `json:"identifier,omitempty"`

	Id      string  `json:"id,omitempty"`

	Promotion      *Promotion  `json:"promotion,omitempty"`

	Downloads      []OrderLineItemDownload  `json:"downloads,omitempty"`

	Parent      *OrderLineItem  `json:"parent,omitempty"`

	PromotionId      string  `json:"promotionId,omitempty"`

	Label      string  `json:"label,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Quantity      float64  `json:"quantity,omitempty"`

	VersionId      string  `json:"versionId,omitempty"`

	ProductVersionId      string  `json:"productVersionId,omitempty"`

	Good      bool  `json:"good,omitempty"`

	ParentId      string  `json:"parentId,omitempty"`

	Payload      interface{}  `json:"payload,omitempty"`

	Removable      bool  `json:"removable,omitempty"`

	Position      float64  `json:"position,omitempty"`

	Children      []OrderLineItem  `json:"children,omitempty"`

	OrderVersionId      string  `json:"orderVersionId,omitempty"`

	ParentVersionId      string  `json:"parentVersionId,omitempty"`

	UnitPrice      float64  `json:"unitPrice,omitempty"`

	Description      string  `json:"description,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

}

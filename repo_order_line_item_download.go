package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type OrderLineItemDownloadRepository struct {
	*GenericRepository[OrderLineItemDownload]
}

func NewOrderLineItemDownloadRepository(client *Client) *OrderLineItemDownloadRepository {
	return &OrderLineItemDownloadRepository{
		GenericRepository: NewGenericRepository[OrderLineItemDownload](client),
	}
}

func (t *OrderLineItemDownloadRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[OrderLineItemDownload], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "order-line-item-download")
}

func (t *OrderLineItemDownloadRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[OrderLineItemDownload], *http.Response, error) {
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

func (t *OrderLineItemDownloadRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "order-line-item-download")
}

func (t *OrderLineItemDownloadRepository) Upsert(ctx ApiContext, entity []OrderLineItemDownload) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "order_line_item_download")
}

func (t *OrderLineItemDownloadRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "order_line_item_download")
}

type OrderLineItemDownload struct {

	AccessGranted      bool  `json:"accessGranted,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Id      string  `json:"id,omitempty"`

	Media      *Media  `json:"media,omitempty"`

	MediaId      string  `json:"mediaId,omitempty"`

	OrderLineItem      *OrderLineItem  `json:"orderLineItem,omitempty"`

	OrderLineItemId      string  `json:"orderLineItemId,omitempty"`

	OrderLineItemVersionId      string  `json:"orderLineItemVersionId,omitempty"`

	Position      float64  `json:"position,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	VersionId      string  `json:"versionId,omitempty"`

}

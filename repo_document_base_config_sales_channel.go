package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type DocumentBaseConfigSalesChannelRepository struct {
	*GenericRepository[DocumentBaseConfigSalesChannel]
}

func NewDocumentBaseConfigSalesChannelRepository(client *Client) *DocumentBaseConfigSalesChannelRepository {
	return &DocumentBaseConfigSalesChannelRepository{
		GenericRepository: NewGenericRepository[DocumentBaseConfigSalesChannel](client),
	}
}

func (t *DocumentBaseConfigSalesChannelRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[DocumentBaseConfigSalesChannel], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "document-base-config-sales-channel")
}

func (t *DocumentBaseConfigSalesChannelRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[DocumentBaseConfigSalesChannel], *http.Response, error) {
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

func (t *DocumentBaseConfigSalesChannelRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "document-base-config-sales-channel")
}

func (t *DocumentBaseConfigSalesChannelRepository) Upsert(ctx ApiContext, entity []DocumentBaseConfigSalesChannel) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "document_base_config_sales_channel")
}

func (t *DocumentBaseConfigSalesChannelRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "document_base_config_sales_channel")
}

type DocumentBaseConfigSalesChannel struct {

	Id      string  `json:"id,omitempty"`

	DocumentBaseConfigId      string  `json:"documentBaseConfigId,omitempty"`

	SalesChannelId      string  `json:"salesChannelId,omitempty"`

	DocumentTypeId      string  `json:"documentTypeId,omitempty"`

	DocumentBaseConfig      *DocumentBaseConfig  `json:"documentBaseConfig,omitempty"`

	SalesChannel      *SalesChannel  `json:"salesChannel,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	DocumentType      *DocumentType  `json:"documentType,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}

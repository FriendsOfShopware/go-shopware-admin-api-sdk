package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type DocumentBaseConfigSalesChannelRepository ClientService

func (t DocumentBaseConfigSalesChannelRepository) Search(ctx ApiContext, criteria Criteria) (*DocumentBaseConfigSalesChannelCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/document-base-config-sales-channel", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(DocumentBaseConfigSalesChannelCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t DocumentBaseConfigSalesChannelRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/document-base-config-sales-channel", criteria)

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

func (t DocumentBaseConfigSalesChannelRepository) Upsert(ctx ApiContext, entity []DocumentBaseConfigSalesChannel) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"document_base_config_sales_channel": {
		Entity:  "document_base_config_sales_channel",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t DocumentBaseConfigSalesChannelRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"document_base_config_sales_channel": {
		Entity:  "document_base_config_sales_channel",
		Action:  "delete",
		Payload: payload,
	}})
}

type DocumentBaseConfigSalesChannel struct {
	SalesChannelId string `json:"salesChannelId,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Id string `json:"id,omitempty"`

	DocumentBaseConfigId string `json:"documentBaseConfigId,omitempty"`

	DocumentTypeId string `json:"documentTypeId,omitempty"`

	DocumentType *DocumentType `json:"documentType,omitempty"`

	DocumentBaseConfig *DocumentBaseConfig `json:"documentBaseConfig,omitempty"`

	SalesChannel *SalesChannel `json:"salesChannel,omitempty"`
}

type DocumentBaseConfigSalesChannelCollection struct {
	EntityCollection

	Data []DocumentBaseConfigSalesChannel `json:"data"`
}

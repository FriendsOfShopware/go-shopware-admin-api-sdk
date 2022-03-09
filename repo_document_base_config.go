package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type DocumentBaseConfigRepository ClientService

func (t DocumentBaseConfigRepository) Search(ctx ApiContext, criteria Criteria) (*DocumentBaseConfigCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/document-base-config", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(DocumentBaseConfigCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t DocumentBaseConfigRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/document-base-config", criteria)

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

func (t DocumentBaseConfigRepository) Upsert(ctx ApiContext, entity []DocumentBaseConfig) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"document_base_config": {
		Entity:  "document_base_config",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t DocumentBaseConfigRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"document_base_config": {
		Entity:  "document_base_config",
		Action:  "delete",
		Payload: payload,
	}})
}

type DocumentBaseConfig struct {
	Id string `json:"id,omitempty"`

	FilenameSuffix string `json:"filenameSuffix,omitempty"`

	DocumentNumber string `json:"documentNumber,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	Logo *Media `json:"logo,omitempty"`

	LogoId string `json:"logoId,omitempty"`

	Config interface{} `json:"config,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	DocumentType *DocumentType `json:"documentType,omitempty"`

	SalesChannels []DocumentBaseConfigSalesChannel `json:"salesChannels,omitempty"`

	DocumentTypeId string `json:"documentTypeId,omitempty"`

	Name string `json:"name,omitempty"`

	FilenamePrefix string `json:"filenamePrefix,omitempty"`

	Global bool `json:"global,omitempty"`
}

type DocumentBaseConfigCollection struct {
	EntityCollection

	Data []DocumentBaseConfig `json:"data"`
}

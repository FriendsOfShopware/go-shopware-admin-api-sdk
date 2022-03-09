package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type DocumentTypeRepository ClientService

func (t DocumentTypeRepository) Search(ctx ApiContext, criteria Criteria) (*DocumentTypeCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/document-type", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(DocumentTypeCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t DocumentTypeRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/document-type", criteria)

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

func (t DocumentTypeRepository) Upsert(ctx ApiContext, entity []DocumentType) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"document_type": {
		Entity:  "document_type",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t DocumentTypeRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"document_type": {
		Entity:  "document_type",
		Action:  "delete",
		Payload: payload,
	}})
}

type DocumentType struct {
	Id string `json:"id,omitempty"`

	TechnicalName string `json:"technicalName,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	Translated interface{} `json:"translated,omitempty"`

	Name string `json:"name,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Translations []DocumentTypeTranslation `json:"translations,omitempty"`

	Documents []Document `json:"documents,omitempty"`

	DocumentBaseConfigs []DocumentBaseConfig `json:"documentBaseConfigs,omitempty"`

	DocumentBaseConfigSalesChannels []DocumentBaseConfigSalesChannel `json:"documentBaseConfigSalesChannels,omitempty"`
}

type DocumentTypeCollection struct {
	EntityCollection

	Data []DocumentType `json:"data"`
}

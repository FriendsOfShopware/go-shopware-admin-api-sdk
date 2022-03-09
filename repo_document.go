package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type DocumentRepository ClientService

func (t DocumentRepository) Search(ctx ApiContext, criteria Criteria) (*DocumentCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/document", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(DocumentCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t DocumentRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/document", criteria)

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

func (t DocumentRepository) Upsert(ctx ApiContext, entity []Document) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"document": {
		Entity:  "document",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t DocumentRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"document": {
		Entity:  "document",
		Action:  "delete",
		Payload: payload,
	}})
}

type Document struct {
	FileType string `json:"fileType,omitempty"`

	Static bool `json:"static,omitempty"`

	DeepLinkCode string `json:"deepLinkCode,omitempty"`

	DocumentType *DocumentType `json:"documentType,omitempty"`

	ReferencedDocument *Document `json:"referencedDocument,omitempty"`

	DocumentMediaFile *Media `json:"documentMediaFile,omitempty"`

	Id string `json:"id,omitempty"`

	DocumentTypeId string `json:"documentTypeId,omitempty"`

	Sent bool `json:"sent,omitempty"`

	Config interface{} `json:"config,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	OrderVersionId string `json:"orderVersionId,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	Order *Order `json:"order,omitempty"`

	DependentDocuments []Document `json:"dependentDocuments,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	ReferencedDocumentId string `json:"referencedDocumentId,omitempty"`

	OrderId string `json:"orderId,omitempty"`

	DocumentMediaFileId string `json:"documentMediaFileId,omitempty"`
}

type DocumentCollection struct {
	EntityCollection

	Data []Document `json:"data"`
}

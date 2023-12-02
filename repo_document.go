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

func (t DocumentRepository) SearchAll(ctx ApiContext, criteria Criteria) (*DocumentCollection, *http.Response, error) {
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

	DeepLinkCode      string  `json:"deepLinkCode,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	DocumentType      *DocumentType  `json:"documentType,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Id      string  `json:"id,omitempty"`

	ReferencedDocumentId      string  `json:"referencedDocumentId,omitempty"`

	OrderVersionId      string  `json:"orderVersionId,omitempty"`

	Sent      bool  `json:"sent,omitempty"`

	Static      bool  `json:"static,omitempty"`

	Order      *Order  `json:"order,omitempty"`

	DependentDocuments      []Document  `json:"dependentDocuments,omitempty"`

	FileType      string  `json:"fileType,omitempty"`

	OrderId      string  `json:"orderId,omitempty"`

	DocumentMediaFileId      string  `json:"documentMediaFileId,omitempty"`

	Config      interface{}  `json:"config,omitempty"`

	DocumentMediaFile      *Media  `json:"documentMediaFile,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	DocumentTypeId      string  `json:"documentTypeId,omitempty"`

	DocumentNumber      string  `json:"documentNumber,omitempty"`

	ReferencedDocument      *Document  `json:"referencedDocument,omitempty"`

}

type DocumentCollection struct {
	EntityCollection

	Data []Document `json:"data"`
}

package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type DocumentRepository struct {
	*GenericRepository[Document]
}

func NewDocumentRepository(client *Client) *DocumentRepository {
	return &DocumentRepository{
		GenericRepository: NewGenericRepository[Document](client),
	}
}

func (t *DocumentRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[Document], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "document")
}

func (t *DocumentRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[Document], *http.Response, error) {
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

func (t *DocumentRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "document")
}

func (t *DocumentRepository) Upsert(ctx ApiContext, entity []Document) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "document")
}

func (t *DocumentRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "document")
}

type Document struct {

	Config      interface{}  `json:"config,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	DeepLinkCode      string  `json:"deepLinkCode,omitempty"`

	DependentDocuments      []Document  `json:"dependentDocuments,omitempty"`

	DocumentA11YMediaFile      *Media  `json:"documentA11yMediaFile,omitempty"`

	DocumentA11YMediaFileId      string  `json:"documentA11yMediaFileId,omitempty"`

	DocumentMediaFile      *Media  `json:"documentMediaFile,omitempty"`

	DocumentMediaFileId      string  `json:"documentMediaFileId,omitempty"`

	DocumentNumber      string  `json:"documentNumber,omitempty"`

	DocumentType      *DocumentType  `json:"documentType,omitempty"`

	DocumentTypeId      string  `json:"documentTypeId,omitempty"`

	FileType      string  `json:"fileType,omitempty"`

	Id      string  `json:"id,omitempty"`

	Order      *Order  `json:"order,omitempty"`

	OrderId      string  `json:"orderId,omitempty"`

	OrderVersionId      string  `json:"orderVersionId,omitempty"`

	ReferencedDocument      *Document  `json:"referencedDocument,omitempty"`

	ReferencedDocumentId      string  `json:"referencedDocumentId,omitempty"`

	Sent      bool  `json:"sent,omitempty"`

	Static      bool  `json:"static,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}

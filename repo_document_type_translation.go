package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type DocumentTypeTranslationRepository struct {
	*GenericRepository[DocumentTypeTranslation]
}

func NewDocumentTypeTranslationRepository(client *Client) *DocumentTypeTranslationRepository {
	return &DocumentTypeTranslationRepository{
		GenericRepository: NewGenericRepository[DocumentTypeTranslation](client),
	}
}

func (t *DocumentTypeTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[DocumentTypeTranslation], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "document-type-translation")
}

func (t *DocumentTypeTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[DocumentTypeTranslation], *http.Response, error) {
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

func (t *DocumentTypeTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "document-type-translation")
}

func (t *DocumentTypeTranslationRepository) Upsert(ctx ApiContext, entity []DocumentTypeTranslation) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "document_type_translation")
}

func (t *DocumentTypeTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "document_type_translation")
}

type DocumentTypeTranslation struct {

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	DocumentType      *DocumentType  `json:"documentType,omitempty"`

	DocumentTypeId      string  `json:"documentTypeId,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	Name      string  `json:"name,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}

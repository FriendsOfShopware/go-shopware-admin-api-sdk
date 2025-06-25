package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type DocumentTypeRepository struct {
	*GenericRepository[DocumentType]
}

func NewDocumentTypeRepository(client *Client) *DocumentTypeRepository {
	return &DocumentTypeRepository{
		GenericRepository: NewGenericRepository[DocumentType](client),
	}
}

func (t *DocumentTypeRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[DocumentType], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "document-type")
}

func (t *DocumentTypeRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[DocumentType], *http.Response, error) {
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

func (t *DocumentTypeRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "document-type")
}

func (t *DocumentTypeRepository) Upsert(ctx ApiContext, entity []DocumentType) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "document_type")
}

func (t *DocumentTypeRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "document_type")
}

type DocumentType struct {

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	DocumentBaseConfigSalesChannels      []DocumentBaseConfigSalesChannel  `json:"documentBaseConfigSalesChannels,omitempty"`

	DocumentBaseConfigs      []DocumentBaseConfig  `json:"documentBaseConfigs,omitempty"`

	Documents      []Document  `json:"documents,omitempty"`

	Id      string  `json:"id,omitempty"`

	Name      string  `json:"name,omitempty"`

	TechnicalName      string  `json:"technicalName,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	Translations      []DocumentTypeTranslation  `json:"translations,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}

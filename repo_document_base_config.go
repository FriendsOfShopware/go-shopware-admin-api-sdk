package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type DocumentBaseConfigRepository struct {
	*GenericRepository[DocumentBaseConfig]
}

func NewDocumentBaseConfigRepository(client *Client) *DocumentBaseConfigRepository {
	return &DocumentBaseConfigRepository{
		GenericRepository: NewGenericRepository[DocumentBaseConfig](client),
	}
}

func (t *DocumentBaseConfigRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[DocumentBaseConfig], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "document-base-config")
}

func (t *DocumentBaseConfigRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[DocumentBaseConfig], *http.Response, error) {
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

func (t *DocumentBaseConfigRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "document-base-config")
}

func (t *DocumentBaseConfigRepository) Upsert(ctx ApiContext, entity []DocumentBaseConfig) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "document_base_config")
}

func (t *DocumentBaseConfigRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "document_base_config")
}

type DocumentBaseConfig struct {

	Id      string  `json:"id,omitempty"`

	FilenameSuffix      string  `json:"filenameSuffix,omitempty"`

	Global      bool  `json:"global,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	LogoId      string  `json:"logoId,omitempty"`

	Name      string  `json:"name,omitempty"`

	DocumentNumber      string  `json:"documentNumber,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	SalesChannels      []DocumentBaseConfigSalesChannel  `json:"salesChannels,omitempty"`

	DocumentTypeId      string  `json:"documentTypeId,omitempty"`

	DocumentType      *DocumentType  `json:"documentType,omitempty"`

	Logo      *Media  `json:"logo,omitempty"`

	FilenamePrefix      string  `json:"filenamePrefix,omitempty"`

	Config      interface{}  `json:"config,omitempty"`

}

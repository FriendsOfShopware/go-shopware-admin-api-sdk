package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type ImportExportProfileRepository struct {
	*GenericRepository[ImportExportProfile]
}

func NewImportExportProfileRepository(client *Client) *ImportExportProfileRepository {
	return &ImportExportProfileRepository{
		GenericRepository: NewGenericRepository[ImportExportProfile](client),
	}
}

func (t *ImportExportProfileRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[ImportExportProfile], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "import-export-profile")
}

func (t *ImportExportProfileRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[ImportExportProfile], *http.Response, error) {
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

func (t *ImportExportProfileRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "import-export-profile")
}

func (t *ImportExportProfileRepository) Upsert(ctx ApiContext, entity []ImportExportProfile) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "import_export_profile")
}

func (t *ImportExportProfileRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "import_export_profile")
}

type ImportExportProfile struct {

	Type      string  `json:"type,omitempty"`

	FileType      string  `json:"fileType,omitempty"`

	Delimiter      string  `json:"delimiter,omitempty"`

	UpdateBy      interface{}  `json:"updateBy,omitempty"`

	TechnicalName      string  `json:"technicalName,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Id      string  `json:"id,omitempty"`

	Label      string  `json:"label,omitempty"`

	Translations      []ImportExportProfileTranslation  `json:"translations,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	Enclosure      string  `json:"enclosure,omitempty"`

	Mapping      interface{}  `json:"mapping,omitempty"`

	Config      interface{}  `json:"config,omitempty"`

	ImportExportLogs      []ImportExportLog  `json:"importExportLogs,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	SystemDefault      bool  `json:"systemDefault,omitempty"`

	SourceEntity      string  `json:"sourceEntity,omitempty"`

	Name      string  `json:"name,omitempty"`

}

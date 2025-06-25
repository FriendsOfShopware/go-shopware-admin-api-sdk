package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type ImportExportFileRepository struct {
	*GenericRepository[ImportExportFile]
}

func NewImportExportFileRepository(client *Client) *ImportExportFileRepository {
	return &ImportExportFileRepository{
		GenericRepository: NewGenericRepository[ImportExportFile](client),
	}
}

func (t *ImportExportFileRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[ImportExportFile], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "import-export-file")
}

func (t *ImportExportFileRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[ImportExportFile], *http.Response, error) {
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

func (t *ImportExportFileRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "import-export-file")
}

func (t *ImportExportFileRepository) Upsert(ctx ApiContext, entity []ImportExportFile) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "import_export_file")
}

func (t *ImportExportFileRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "import_export_file")
}

type ImportExportFile struct {

	AccessToken      string  `json:"accessToken,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	ExpireDate      time.Time  `json:"expireDate,omitempty"`

	Id      string  `json:"id,omitempty"`

	Log      *ImportExportLog  `json:"log,omitempty"`

	OriginalName      string  `json:"originalName,omitempty"`

	Path      string  `json:"path,omitempty"`

	Size      float64  `json:"size,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}

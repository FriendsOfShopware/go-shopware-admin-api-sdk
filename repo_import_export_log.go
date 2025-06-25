package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type ImportExportLogRepository struct {
	*GenericRepository[ImportExportLog]
}

func NewImportExportLogRepository(client *Client) *ImportExportLogRepository {
	return &ImportExportLogRepository{
		GenericRepository: NewGenericRepository[ImportExportLog](client),
	}
}

func (t *ImportExportLogRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[ImportExportLog], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "import-export-log")
}

func (t *ImportExportLogRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[ImportExportLog], *http.Response, error) {
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

func (t *ImportExportLogRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "import-export-log")
}

func (t *ImportExportLogRepository) Upsert(ctx ApiContext, entity []ImportExportLog) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "import_export_log")
}

func (t *ImportExportLogRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "import_export_log")
}

type ImportExportLog struct {

	Activity      string  `json:"activity,omitempty"`

	Config      interface{}  `json:"config,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	FailedImportLog      *ImportExportLog  `json:"failedImportLog,omitempty"`

	File      *ImportExportFile  `json:"file,omitempty"`

	FileId      string  `json:"fileId,omitempty"`

	Id      string  `json:"id,omitempty"`

	InvalidRecordsLog      *ImportExportLog  `json:"invalidRecordsLog,omitempty"`

	InvalidRecordsLogId      string  `json:"invalidRecordsLogId,omitempty"`

	Profile      *ImportExportProfile  `json:"profile,omitempty"`

	ProfileId      string  `json:"profileId,omitempty"`

	ProfileName      string  `json:"profileName,omitempty"`

	Records      float64  `json:"records,omitempty"`

	Result      interface{}  `json:"result,omitempty"`

	State      string  `json:"state,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	User      *User  `json:"user,omitempty"`

	UserId      string  `json:"userId,omitempty"`

	Username      string  `json:"username,omitempty"`

}

package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type ImportExportLogRepository ClientService

func (t ImportExportLogRepository) Search(ctx ApiContext, criteria Criteria) (*ImportExportLogCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/import-export-log", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(ImportExportLogCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ImportExportLogRepository) SearchAll(ctx ApiContext, criteria Criteria) (*ImportExportLogCollection, *http.Response, error) {
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

func (t ImportExportLogRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/import-export-log", criteria)

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

func (t ImportExportLogRepository) Upsert(ctx ApiContext, entity []ImportExportLog) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"import_export_log": {
		Entity:  "import_export_log",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t ImportExportLogRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"import_export_log": {
		Entity:  "import_export_log",
		Action:  "delete",
		Payload: payload,
	}})
}

type ImportExportLog struct {
	Profile *ImportExportProfile `json:"profile,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	State string `json:"state,omitempty"`

	User *User `json:"user,omitempty"`

	InvalidRecordsLogId string `json:"invalidRecordsLogId,omitempty"`

	Username string `json:"username,omitempty"`

	FailedImportLog *ImportExportLog `json:"failedImportLog,omitempty"`

	Id string `json:"id,omitempty"`

	UserId string `json:"userId,omitempty"`

	ProfileName string `json:"profileName,omitempty"`

	Result interface{} `json:"result,omitempty"`

	File *ImportExportFile `json:"file,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Records float64 `json:"records,omitempty"`

	ProfileId string `json:"profileId,omitempty"`

	Config interface{} `json:"config,omitempty"`

	InvalidRecordsLog *ImportExportLog `json:"invalidRecordsLog,omitempty"`

	Activity string `json:"activity,omitempty"`

	FileId string `json:"fileId,omitempty"`
}

type ImportExportLogCollection struct {
	EntityCollection

	Data []ImportExportLog `json:"data"`
}

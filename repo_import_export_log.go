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
	Id string `json:"id,omitempty"`

	UserId string `json:"userId,omitempty"`

	ProfileName string `json:"profileName,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	FileId string `json:"fileId,omitempty"`

	InvalidRecordsLogId string `json:"invalidRecordsLogId,omitempty"`

	User *User `json:"user,omitempty"`

	Profile *ImportExportProfile `json:"profile,omitempty"`

	File *ImportExportFile `json:"file,omitempty"`

	Activity string `json:"activity,omitempty"`

	State string `json:"state,omitempty"`

	Records float64 `json:"records,omitempty"`

	ProfileId string `json:"profileId,omitempty"`

	Config interface{} `json:"config,omitempty"`

	Result interface{} `json:"result,omitempty"`

	InvalidRecordsLog *ImportExportLog `json:"invalidRecordsLog,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Username string `json:"username,omitempty"`

	FailedImportLog *ImportExportLog `json:"failedImportLog,omitempty"`
}

type ImportExportLogCollection struct {
	EntityCollection

	Data []ImportExportLog `json:"data"`
}

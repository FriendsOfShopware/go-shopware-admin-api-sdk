package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type ImportExportFileRepository ClientService

func (t ImportExportFileRepository) Search(ctx ApiContext, criteria Criteria) (*ImportExportFileCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/import-export-file", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(ImportExportFileCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ImportExportFileRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/import-export-file", criteria)

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

func (t ImportExportFileRepository) Upsert(ctx ApiContext, entity []ImportExportFile) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"import_export_file": {
		Entity:  "import_export_file",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t ImportExportFileRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"import_export_file": {
		Entity:  "import_export_file",
		Action:  "delete",
		Payload: payload,
	}})
}

type ImportExportFile struct {
	Path string `json:"path,omitempty"`

	ExpireDate time.Time `json:"expireDate,omitempty"`

	AccessToken string `json:"accessToken,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	OriginalName string `json:"originalName,omitempty"`

	Size float64 `json:"size,omitempty"`

	Log *ImportExportLog `json:"log,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Id string `json:"id,omitempty"`
}

type ImportExportFileCollection struct {
	EntityCollection

	Data []ImportExportFile `json:"data"`
}

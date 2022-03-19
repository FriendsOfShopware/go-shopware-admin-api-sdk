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

func (t ImportExportFileRepository) SearchAll(ctx ApiContext, criteria Criteria) (*ImportExportFileCollection, *http.Response, error) {
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
	AccessToken string `json:"accessToken,omitempty"`

	ExpireDate time.Time `json:"expireDate,omitempty"`

	Size float64 `json:"size,omitempty"`

	Log *ImportExportLog `json:"log,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Id string `json:"id,omitempty"`

	OriginalName string `json:"originalName,omitempty"`

	Path string `json:"path,omitempty"`
}

type ImportExportFileCollection struct {
	EntityCollection

	Data []ImportExportFile `json:"data"`
}

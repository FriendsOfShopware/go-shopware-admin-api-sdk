package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type ImportExportProfileRepository ClientService

func (t ImportExportProfileRepository) Search(ctx ApiContext, criteria Criteria) (*ImportExportProfileCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/import-export-profile", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(ImportExportProfileCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ImportExportProfileRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/import-export-profile", criteria)

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

func (t ImportExportProfileRepository) Upsert(ctx ApiContext, entity []ImportExportProfile) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"import_export_profile": {
		Entity:  "import_export_profile",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t ImportExportProfileRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"import_export_profile": {
		Entity:  "import_export_profile",
		Action:  "delete",
		Payload: payload,
	}})
}

type ImportExportProfile struct {
	Type string `json:"type,omitempty"`

	ImportExportLogs []ImportExportLog `json:"importExportLogs,omitempty"`

	Translations []ImportExportProfileTranslation `json:"translations,omitempty"`

	Translated interface{} `json:"translated,omitempty"`

	Name string `json:"name,omitempty"`

	Label string `json:"label,omitempty"`

	Id string `json:"id,omitempty"`

	Enclosure string `json:"enclosure,omitempty"`

	Config interface{} `json:"config,omitempty"`

	FileType string `json:"fileType,omitempty"`

	Delimiter string `json:"delimiter,omitempty"`

	Mapping interface{} `json:"mapping,omitempty"`

	UpdateBy interface{} `json:"updateBy,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	SystemDefault bool `json:"systemDefault,omitempty"`

	SourceEntity string `json:"sourceEntity,omitempty"`
}

type ImportExportProfileCollection struct {
	EntityCollection

	Data []ImportExportProfile `json:"data"`
}

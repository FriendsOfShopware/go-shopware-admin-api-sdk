package go_shopware_admin_sdk

import (
	"net/http"

	"time"
)

type ImportExportProfileTranslationRepository ClientService

func (t ImportExportProfileTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*ImportExportProfileTranslationCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/import-export-profile-translation", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(ImportExportProfileTranslationCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ImportExportProfileTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*ImportExportProfileTranslationCollection, *http.Response, error) {
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

func (t ImportExportProfileTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/import-export-profile-translation", criteria)

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

func (t ImportExportProfileTranslationRepository) Upsert(ctx ApiContext, entity []ImportExportProfileTranslation) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"import_export_profile_translation": {
		Entity:  "import_export_profile_translation",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t ImportExportProfileTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"import_export_profile_translation": {
		Entity:  "import_export_profile_translation",
		Action:  "delete",
		Payload: payload,
	}})
}

type ImportExportProfileTranslation struct {
	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	ImportExportProfileId string `json:"importExportProfileId,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	ImportExportProfile *ImportExportProfile `json:"importExportProfile,omitempty"`

	Language *Language `json:"language,omitempty"`

	Label string `json:"label,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`
}

type ImportExportProfileTranslationCollection struct {
	EntityCollection

	Data []ImportExportProfileTranslation `json:"data"`
}

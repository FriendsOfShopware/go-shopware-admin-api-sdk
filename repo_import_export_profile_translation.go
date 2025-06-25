package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type ImportExportProfileTranslationRepository struct {
	*GenericRepository[ImportExportProfileTranslation]
}

func NewImportExportProfileTranslationRepository(client *Client) *ImportExportProfileTranslationRepository {
	return &ImportExportProfileTranslationRepository{
		GenericRepository: NewGenericRepository[ImportExportProfileTranslation](client),
	}
}

func (t *ImportExportProfileTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[ImportExportProfileTranslation], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "import-export-profile-translation")
}

func (t *ImportExportProfileTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[ImportExportProfileTranslation], *http.Response, error) {
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

func (t *ImportExportProfileTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "import-export-profile-translation")
}

func (t *ImportExportProfileTranslationRepository) Upsert(ctx ApiContext, entity []ImportExportProfileTranslation) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "import_export_profile_translation")
}

func (t *ImportExportProfileTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "import_export_profile_translation")
}

type ImportExportProfileTranslation struct {

	Label      string  `json:"label,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	ImportExportProfileId      string  `json:"importExportProfileId,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	ImportExportProfile      *ImportExportProfile  `json:"importExportProfile,omitempty"`

	Language      *Language  `json:"language,omitempty"`

}

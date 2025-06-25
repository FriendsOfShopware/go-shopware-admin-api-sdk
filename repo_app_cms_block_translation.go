package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type AppCmsBlockTranslationRepository struct {
	*GenericRepository[AppCmsBlockTranslation]
}

func NewAppCmsBlockTranslationRepository(client *Client) *AppCmsBlockTranslationRepository {
	return &AppCmsBlockTranslationRepository{
		GenericRepository: NewGenericRepository[AppCmsBlockTranslation](client),
	}
}

func (t *AppCmsBlockTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[AppCmsBlockTranslation], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "app-cms-block-translation")
}

func (t *AppCmsBlockTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[AppCmsBlockTranslation], *http.Response, error) {
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

func (t *AppCmsBlockTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "app-cms-block-translation")
}

func (t *AppCmsBlockTranslationRepository) Upsert(ctx ApiContext, entity []AppCmsBlockTranslation) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "app_cms_block_translation")
}

func (t *AppCmsBlockTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "app_cms_block_translation")
}

type AppCmsBlockTranslation struct {

	AppCmsBlock      *AppCmsBlock  `json:"appCmsBlock,omitempty"`

	AppCmsBlockId      string  `json:"appCmsBlockId,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Label      string  `json:"label,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}

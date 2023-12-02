package go_shopware_admin_sdk

import (
	"net/http"

	"time"
)

type AppCmsBlockTranslationRepository ClientService

func (t AppCmsBlockTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*AppCmsBlockTranslationCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/app-cms-block-translation", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(AppCmsBlockTranslationCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t AppCmsBlockTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*AppCmsBlockTranslationCollection, *http.Response, error) {
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

func (t AppCmsBlockTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/app-cms-block-translation", criteria)

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

func (t AppCmsBlockTranslationRepository) Upsert(ctx ApiContext, entity []AppCmsBlockTranslation) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"app_cms_block_translation": {
		Entity:  "app_cms_block_translation",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t AppCmsBlockTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"app_cms_block_translation": {
		Entity:  "app_cms_block_translation",
		Action:  "delete",
		Payload: payload,
	}})
}

type AppCmsBlockTranslation struct {
	AppCmsBlockId string `json:"appCmsBlockId,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	AppCmsBlock *AppCmsBlock `json:"appCmsBlock,omitempty"`

	Language *Language `json:"language,omitempty"`

	Label string `json:"label,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type AppCmsBlockTranslationCollection struct {
	EntityCollection

	Data []AppCmsBlockTranslation `json:"data"`
}

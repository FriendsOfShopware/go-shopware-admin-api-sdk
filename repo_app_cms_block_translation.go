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
	LanguageId string `json:"languageId,omitempty"`

	AppCmsBlock *AppCmsBlock `json:"appCmsBlock,omitempty"`

	Language *Language `json:"language,omitempty"`

	Label string `json:"label,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	AppCmsBlockId string `json:"appCmsBlockId,omitempty"`
}

type AppCmsBlockTranslationCollection struct {
	EntityCollection

	Data []AppCmsBlockTranslation `json:"data"`
}

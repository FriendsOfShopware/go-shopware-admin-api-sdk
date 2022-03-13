package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type DocumentTypeTranslationRepository ClientService

func (t DocumentTypeTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*DocumentTypeTranslationCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/document-type-translation", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(DocumentTypeTranslationCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t DocumentTypeTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/document-type-translation", criteria)

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

func (t DocumentTypeTranslationRepository) Upsert(ctx ApiContext, entity []DocumentTypeTranslation) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"document_type_translation": {
		Entity:  "document_type_translation",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t DocumentTypeTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"document_type_translation": {
		Entity:  "document_type_translation",
		Action:  "delete",
		Payload: payload,
	}})
}

type DocumentTypeTranslation struct {
	Language *Language `json:"language,omitempty"`

	Name string `json:"name,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	DocumentTypeId string `json:"documentTypeId,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	DocumentType *DocumentType `json:"documentType,omitempty"`
}

type DocumentTypeTranslationCollection struct {
	EntityCollection

	Data []DocumentTypeTranslation `json:"data"`
}
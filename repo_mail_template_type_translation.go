package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type MailTemplateTypeTranslationRepository ClientService

func (t MailTemplateTypeTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*MailTemplateTypeTranslationCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/mail-template-type-translation", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(MailTemplateTypeTranslationCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t MailTemplateTypeTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/mail-template-type-translation", criteria)

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

func (t MailTemplateTypeTranslationRepository) Upsert(ctx ApiContext, entity []MailTemplateTypeTranslation) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"mail_template_type_translation": {
		Entity:  "mail_template_type_translation",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t MailTemplateTypeTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"mail_template_type_translation": {
		Entity:  "mail_template_type_translation",
		Action:  "delete",
		Payload: payload,
	}})
}

type MailTemplateTypeTranslation struct {
	Language *Language `json:"language,omitempty"`

	Name string `json:"name,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	MailTemplateTypeId string `json:"mailTemplateTypeId,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	MailTemplateType *MailTemplateType `json:"mailTemplateType,omitempty"`
}

type MailTemplateTypeTranslationCollection struct {
	EntityCollection

	Data []MailTemplateTypeTranslation `json:"data"`
}

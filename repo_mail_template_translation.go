package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type MailTemplateTranslationRepository ClientService

func (t MailTemplateTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*MailTemplateTranslationCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/mail-template-translation", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(MailTemplateTranslationCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t MailTemplateTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/mail-template-translation", criteria)

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

func (t MailTemplateTranslationRepository) Upsert(ctx ApiContext, entity []MailTemplateTranslation) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"mail_template_translation": {
		Entity:  "mail_template_translation",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t MailTemplateTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"mail_template_translation": {
		Entity:  "mail_template_translation",
		Action:  "delete",
		Payload: payload,
	}})
}

type MailTemplateTranslation struct {
	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	Language *Language `json:"language,omitempty"`

	SenderName string `json:"senderName,omitempty"`

	Description string `json:"description,omitempty"`

	Subject string `json:"subject,omitempty"`

	ContentHtml string `json:"contentHtml,omitempty"`

	ContentPlain string `json:"contentPlain,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	MailTemplateId string `json:"mailTemplateId,omitempty"`

	MailTemplate *MailTemplate `json:"mailTemplate,omitempty"`
}

type MailTemplateTranslationCollection struct {
	EntityCollection

	Data []MailTemplateTranslation `json:"data"`
}

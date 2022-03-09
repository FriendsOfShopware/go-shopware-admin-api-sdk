package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type MailTemplateRepository ClientService

func (t MailTemplateRepository) Search(ctx ApiContext, criteria Criteria) (*MailTemplateCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/mail-template", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(MailTemplateCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t MailTemplateRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/mail-template", criteria)

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

func (t MailTemplateRepository) Upsert(ctx ApiContext, entity []MailTemplate) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"mail_template": {
		Entity:  "mail_template",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t MailTemplateRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"mail_template": {
		Entity:  "mail_template",
		Action:  "delete",
		Payload: payload,
	}})
}

type MailTemplate struct {
	Description string `json:"description,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	MailTemplateType *MailTemplateType `json:"mailTemplateType,omitempty"`

	Translated interface{} `json:"translated,omitempty"`

	SystemDefault bool `json:"systemDefault,omitempty"`

	ContentPlain string `json:"contentPlain,omitempty"`

	SenderName string `json:"senderName,omitempty"`

	Media []MailTemplateMedia `json:"media,omitempty"`

	Id string `json:"id,omitempty"`

	Subject string `json:"subject,omitempty"`

	ContentHtml string `json:"contentHtml,omitempty"`

	Translations []MailTemplateTranslation `json:"translations,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	MailTemplateTypeId string `json:"mailTemplateTypeId,omitempty"`
}

type MailTemplateCollection struct {
	EntityCollection

	Data []MailTemplate `json:"data"`
}

package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type MailTemplateTypeRepository ClientService

func (t MailTemplateTypeRepository) Search(ctx ApiContext, criteria Criteria) (*MailTemplateTypeCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/mail-template-type", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(MailTemplateTypeCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t MailTemplateTypeRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/mail-template-type", criteria)

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

func (t MailTemplateTypeRepository) Upsert(ctx ApiContext, entity []MailTemplateType) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"mail_template_type": {
		Entity:  "mail_template_type",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t MailTemplateTypeRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"mail_template_type": {
		Entity:  "mail_template_type",
		Action:  "delete",
		Payload: payload,
	}})
}

type MailTemplateType struct {
	Translations []MailTemplateTypeTranslation `json:"translations,omitempty"`

	TechnicalName string `json:"technicalName,omitempty"`

	AvailableEntities interface{} `json:"availableEntities,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	MailTemplates []MailTemplate `json:"mailTemplates,omitempty"`

	TemplateData interface{} `json:"templateData,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Translated interface{} `json:"translated,omitempty"`

	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`
}

type MailTemplateTypeCollection struct {
	EntityCollection

	Data []MailTemplateType `json:"data"`
}

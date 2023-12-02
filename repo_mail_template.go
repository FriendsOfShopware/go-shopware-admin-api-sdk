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

func (t MailTemplateRepository) SearchAll(ctx ApiContext, criteria Criteria) (*MailTemplateCollection, *http.Response, error) {
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

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Translations      []MailTemplateTranslation  `json:"translations,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	Id      string  `json:"id,omitempty"`

	SystemDefault      bool  `json:"systemDefault,omitempty"`

	Description      string  `json:"description,omitempty"`

	Subject      string  `json:"subject,omitempty"`

	ContentPlain      string  `json:"contentPlain,omitempty"`

	MailTemplateType      *MailTemplateType  `json:"mailTemplateType,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	MailTemplateTypeId      string  `json:"mailTemplateTypeId,omitempty"`

	SenderName      string  `json:"senderName,omitempty"`

	ContentHtml      string  `json:"contentHtml,omitempty"`

	Media      []MailTemplateMedia  `json:"media,omitempty"`

}

type MailTemplateCollection struct {
	EntityCollection

	Data []MailTemplate `json:"data"`
}

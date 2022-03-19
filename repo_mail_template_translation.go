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

func (t MailTemplateTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*MailTemplateTranslationCollection, *http.Response, error) {
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
	SenderName string `json:"senderName,omitempty"`

	Description string `json:"description,omitempty"`

	Subject string `json:"subject,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	ContentHtml string `json:"contentHtml,omitempty"`

	ContentPlain string `json:"contentPlain,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	MailTemplateId string `json:"mailTemplateId,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	MailTemplate *MailTemplate `json:"mailTemplate,omitempty"`

	Language *Language `json:"language,omitempty"`
}

type MailTemplateTranslationCollection struct {
	EntityCollection

	Data []MailTemplateTranslation `json:"data"`
}

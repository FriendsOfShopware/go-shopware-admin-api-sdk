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

func (t MailTemplateTypeTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*MailTemplateTypeTranslationCollection, *http.Response, error) {
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
	Name string `json:"name,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	MailTemplateTypeId string `json:"mailTemplateTypeId,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	MailTemplateType *MailTemplateType `json:"mailTemplateType,omitempty"`

	Language *Language `json:"language,omitempty"`
}

type MailTemplateTypeTranslationCollection struct {
	EntityCollection

	Data []MailTemplateTypeTranslation `json:"data"`
}

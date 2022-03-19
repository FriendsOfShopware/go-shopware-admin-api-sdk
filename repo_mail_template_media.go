package go_shopware_admin_sdk

import (
	"net/http"
)

type MailTemplateMediaRepository ClientService

func (t MailTemplateMediaRepository) Search(ctx ApiContext, criteria Criteria) (*MailTemplateMediaCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/mail-template-media", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(MailTemplateMediaCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t MailTemplateMediaRepository) SearchAll(ctx ApiContext, criteria Criteria) (*MailTemplateMediaCollection, *http.Response, error) {
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

func (t MailTemplateMediaRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/mail-template-media", criteria)

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

func (t MailTemplateMediaRepository) Upsert(ctx ApiContext, entity []MailTemplateMedia) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"mail_template_media": {
		Entity:  "mail_template_media",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t MailTemplateMediaRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"mail_template_media": {
		Entity:  "mail_template_media",
		Action:  "delete",
		Payload: payload,
	}})
}

type MailTemplateMedia struct {
	MediaId string `json:"mediaId,omitempty"`

	Position float64 `json:"position,omitempty"`

	MailTemplate *MailTemplate `json:"mailTemplate,omitempty"`

	Media *Media `json:"media,omitempty"`

	Id string `json:"id,omitempty"`

	MailTemplateId string `json:"mailTemplateId,omitempty"`

	LanguageId string `json:"languageId,omitempty"`
}

type MailTemplateMediaCollection struct {
	EntityCollection

	Data []MailTemplateMedia `json:"data"`
}

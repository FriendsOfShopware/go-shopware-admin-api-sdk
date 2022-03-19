package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type MailHeaderFooterTranslationRepository ClientService

func (t MailHeaderFooterTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*MailHeaderFooterTranslationCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/mail-header-footer-translation", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(MailHeaderFooterTranslationCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t MailHeaderFooterTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*MailHeaderFooterTranslationCollection, *http.Response, error) {
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

func (t MailHeaderFooterTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/mail-header-footer-translation", criteria)

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

func (t MailHeaderFooterTranslationRepository) Upsert(ctx ApiContext, entity []MailHeaderFooterTranslation) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"mail_header_footer_translation": {
		Entity:  "mail_header_footer_translation",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t MailHeaderFooterTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"mail_header_footer_translation": {
		Entity:  "mail_header_footer_translation",
		Action:  "delete",
		Payload: payload,
	}})
}

type MailHeaderFooterTranslation struct {
	Description string `json:"description,omitempty"`

	HeaderPlain string `json:"headerPlain,omitempty"`

	FooterHtml string `json:"footerHtml,omitempty"`

	FooterPlain string `json:"footerPlain,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	MailHeaderFooterId string `json:"mailHeaderFooterId,omitempty"`

	MailHeaderFooter *MailHeaderFooter `json:"mailHeaderFooter,omitempty"`

	Name string `json:"name,omitempty"`

	Language *Language `json:"language,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	HeaderHtml string `json:"headerHtml,omitempty"`
}

type MailHeaderFooterTranslationCollection struct {
	EntityCollection

	Data []MailHeaderFooterTranslation `json:"data"`
}

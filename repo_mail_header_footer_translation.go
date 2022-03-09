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
	MailHeaderFooterId string `json:"mailHeaderFooterId,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	FooterHtml string `json:"footerHtml,omitempty"`

	FooterPlain string `json:"footerPlain,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	MailHeaderFooter *MailHeaderFooter `json:"mailHeaderFooter,omitempty"`

	Language *Language `json:"language,omitempty"`

	Name string `json:"name,omitempty"`

	Description string `json:"description,omitempty"`

	HeaderHtml string `json:"headerHtml,omitempty"`

	HeaderPlain string `json:"headerPlain,omitempty"`
}

type MailHeaderFooterTranslationCollection struct {
	EntityCollection

	Data []MailHeaderFooterTranslation `json:"data"`
}

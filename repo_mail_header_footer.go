package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type MailHeaderFooterRepository ClientService

func (t MailHeaderFooterRepository) Search(ctx ApiContext, criteria Criteria) (*MailHeaderFooterCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/mail-header-footer", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(MailHeaderFooterCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t MailHeaderFooterRepository) SearchAll(ctx ApiContext, criteria Criteria) (*MailHeaderFooterCollection, *http.Response, error) {
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

func (t MailHeaderFooterRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/mail-header-footer", criteria)

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

func (t MailHeaderFooterRepository) Upsert(ctx ApiContext, entity []MailHeaderFooter) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"mail_header_footer": {
		Entity:  "mail_header_footer",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t MailHeaderFooterRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"mail_header_footer": {
		Entity:  "mail_header_footer",
		Action:  "delete",
		Payload: payload,
	}})
}

type MailHeaderFooter struct {
	HeaderPlain string `json:"headerPlain,omitempty"`

	FooterHtml string `json:"footerHtml,omitempty"`

	Translations []MailHeaderFooterTranslation `json:"translations,omitempty"`

	SalesChannels []SalesChannel `json:"salesChannels,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	HeaderHtml string `json:"headerHtml,omitempty"`

	SystemDefault bool `json:"systemDefault,omitempty"`

	Name string `json:"name,omitempty"`

	Description string `json:"description,omitempty"`

	FooterPlain string `json:"footerPlain,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Translated interface{} `json:"translated,omitempty"`

	Id string `json:"id,omitempty"`
}

type MailHeaderFooterCollection struct {
	EntityCollection

	Data []MailHeaderFooter `json:"data"`
}

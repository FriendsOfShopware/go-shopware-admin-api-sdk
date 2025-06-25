package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type MailTemplateRepository struct {
	*GenericRepository[MailTemplate]
}

func NewMailTemplateRepository(client *Client) *MailTemplateRepository {
	return &MailTemplateRepository{
		GenericRepository: NewGenericRepository[MailTemplate](client),
	}
}

func (t *MailTemplateRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[MailTemplate], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "mail-template")
}

func (t *MailTemplateRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[MailTemplate], *http.Response, error) {
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

func (t *MailTemplateRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "mail-template")
}

func (t *MailTemplateRepository) Upsert(ctx ApiContext, entity []MailTemplate) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "mail_template")
}

func (t *MailTemplateRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "mail_template")
}

type MailTemplate struct {

	ContentHtml      string  `json:"contentHtml,omitempty"`

	ContentPlain      string  `json:"contentPlain,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Description      string  `json:"description,omitempty"`

	Id      string  `json:"id,omitempty"`

	MailTemplateType      *MailTemplateType  `json:"mailTemplateType,omitempty"`

	MailTemplateTypeId      string  `json:"mailTemplateTypeId,omitempty"`

	Media      []MailTemplateMedia  `json:"media,omitempty"`

	SenderName      string  `json:"senderName,omitempty"`

	Subject      string  `json:"subject,omitempty"`

	SystemDefault      bool  `json:"systemDefault,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	Translations      []MailTemplateTranslation  `json:"translations,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}

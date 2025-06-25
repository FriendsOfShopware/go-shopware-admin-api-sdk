package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type MailTemplateTranslationRepository struct {
	*GenericRepository[MailTemplateTranslation]
}

func NewMailTemplateTranslationRepository(client *Client) *MailTemplateTranslationRepository {
	return &MailTemplateTranslationRepository{
		GenericRepository: NewGenericRepository[MailTemplateTranslation](client),
	}
}

func (t *MailTemplateTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[MailTemplateTranslation], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "mail-template-translation")
}

func (t *MailTemplateTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[MailTemplateTranslation], *http.Response, error) {
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

func (t *MailTemplateTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "mail-template-translation")
}

func (t *MailTemplateTranslationRepository) Upsert(ctx ApiContext, entity []MailTemplateTranslation) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "mail_template_translation")
}

func (t *MailTemplateTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "mail_template_translation")
}

type MailTemplateTranslation struct {

	Language      *Language  `json:"language,omitempty"`

	SenderName      string  `json:"senderName,omitempty"`

	Subject      string  `json:"subject,omitempty"`

	ContentPlain      string  `json:"contentPlain,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	MailTemplateId      string  `json:"mailTemplateId,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	Description      string  `json:"description,omitempty"`

	ContentHtml      string  `json:"contentHtml,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	MailTemplate      *MailTemplate  `json:"mailTemplate,omitempty"`

}

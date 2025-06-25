package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type MailTemplateTypeTranslationRepository struct {
	*GenericRepository[MailTemplateTypeTranslation]
}

func NewMailTemplateTypeTranslationRepository(client *Client) *MailTemplateTypeTranslationRepository {
	return &MailTemplateTypeTranslationRepository{
		GenericRepository: NewGenericRepository[MailTemplateTypeTranslation](client),
	}
}

func (t *MailTemplateTypeTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[MailTemplateTypeTranslation], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "mail-template-type-translation")
}

func (t *MailTemplateTypeTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[MailTemplateTypeTranslation], *http.Response, error) {
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

func (t *MailTemplateTypeTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "mail-template-type-translation")
}

func (t *MailTemplateTypeTranslationRepository) Upsert(ctx ApiContext, entity []MailTemplateTypeTranslation) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "mail_template_type_translation")
}

func (t *MailTemplateTypeTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "mail_template_type_translation")
}

type MailTemplateTypeTranslation struct {

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	MailTemplateType      *MailTemplateType  `json:"mailTemplateType,omitempty"`

	MailTemplateTypeId      string  `json:"mailTemplateTypeId,omitempty"`

	Name      string  `json:"name,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}

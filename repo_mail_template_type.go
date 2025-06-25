package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type MailTemplateTypeRepository struct {
	*GenericRepository[MailTemplateType]
}

func NewMailTemplateTypeRepository(client *Client) *MailTemplateTypeRepository {
	return &MailTemplateTypeRepository{
		GenericRepository: NewGenericRepository[MailTemplateType](client),
	}
}

func (t *MailTemplateTypeRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[MailTemplateType], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "mail-template-type")
}

func (t *MailTemplateTypeRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[MailTemplateType], *http.Response, error) {
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

func (t *MailTemplateTypeRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "mail-template-type")
}

func (t *MailTemplateTypeRepository) Upsert(ctx ApiContext, entity []MailTemplateType) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "mail_template_type")
}

func (t *MailTemplateTypeRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "mail_template_type")
}

type MailTemplateType struct {

	AvailableEntities      interface{}  `json:"availableEntities,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Id      string  `json:"id,omitempty"`

	MailTemplates      []MailTemplate  `json:"mailTemplates,omitempty"`

	Name      string  `json:"name,omitempty"`

	TechnicalName      string  `json:"technicalName,omitempty"`

	TemplateData      interface{}  `json:"templateData,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	Translations      []MailTemplateTypeTranslation  `json:"translations,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}

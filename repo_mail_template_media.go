package go_shopware_admin_sdk

import (
	"net/http"

)

type MailTemplateMediaRepository struct {
	*GenericRepository[MailTemplateMedia]
}

func NewMailTemplateMediaRepository(client *Client) *MailTemplateMediaRepository {
	return &MailTemplateMediaRepository{
		GenericRepository: NewGenericRepository[MailTemplateMedia](client),
	}
}

func (t *MailTemplateMediaRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[MailTemplateMedia], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "mail-template-media")
}

func (t *MailTemplateMediaRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[MailTemplateMedia], *http.Response, error) {
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

func (t *MailTemplateMediaRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "mail-template-media")
}

func (t *MailTemplateMediaRepository) Upsert(ctx ApiContext, entity []MailTemplateMedia) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "mail_template_media")
}

func (t *MailTemplateMediaRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "mail_template_media")
}

type MailTemplateMedia struct {

	Id      string  `json:"id,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	MailTemplate      *MailTemplate  `json:"mailTemplate,omitempty"`

	MailTemplateId      string  `json:"mailTemplateId,omitempty"`

	Media      *Media  `json:"media,omitempty"`

	MediaId      string  `json:"mediaId,omitempty"`

	Position      float64  `json:"position,omitempty"`

}

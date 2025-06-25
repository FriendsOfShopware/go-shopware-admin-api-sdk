package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type MailHeaderFooterTranslationRepository struct {
	*GenericRepository[MailHeaderFooterTranslation]
}

func NewMailHeaderFooterTranslationRepository(client *Client) *MailHeaderFooterTranslationRepository {
	return &MailHeaderFooterTranslationRepository{
		GenericRepository: NewGenericRepository[MailHeaderFooterTranslation](client),
	}
}

func (t *MailHeaderFooterTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[MailHeaderFooterTranslation], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "mail-header-footer-translation")
}

func (t *MailHeaderFooterTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[MailHeaderFooterTranslation], *http.Response, error) {
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

func (t *MailHeaderFooterTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "mail-header-footer-translation")
}

func (t *MailHeaderFooterTranslationRepository) Upsert(ctx ApiContext, entity []MailHeaderFooterTranslation) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "mail_header_footer_translation")
}

func (t *MailHeaderFooterTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "mail_header_footer_translation")
}

type MailHeaderFooterTranslation struct {

	HeaderHtml      string  `json:"headerHtml,omitempty"`

	FooterPlain      string  `json:"footerPlain,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	Description      string  `json:"description,omitempty"`

	HeaderPlain      string  `json:"headerPlain,omitempty"`

	FooterHtml      string  `json:"footerHtml,omitempty"`

	MailHeaderFooterId      string  `json:"mailHeaderFooterId,omitempty"`

	MailHeaderFooter      *MailHeaderFooter  `json:"mailHeaderFooter,omitempty"`

	Name      string  `json:"name,omitempty"`

}

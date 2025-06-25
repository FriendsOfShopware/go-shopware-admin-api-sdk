package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type MailHeaderFooterRepository struct {
	*GenericRepository[MailHeaderFooter]
}

func NewMailHeaderFooterRepository(client *Client) *MailHeaderFooterRepository {
	return &MailHeaderFooterRepository{
		GenericRepository: NewGenericRepository[MailHeaderFooter](client),
	}
}

func (t *MailHeaderFooterRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[MailHeaderFooter], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "mail-header-footer")
}

func (t *MailHeaderFooterRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[MailHeaderFooter], *http.Response, error) {
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

func (t *MailHeaderFooterRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "mail-header-footer")
}

func (t *MailHeaderFooterRepository) Upsert(ctx ApiContext, entity []MailHeaderFooter) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "mail_header_footer")
}

func (t *MailHeaderFooterRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "mail_header_footer")
}

type MailHeaderFooter struct {

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	HeaderHtml      string  `json:"headerHtml,omitempty"`

	HeaderPlain      string  `json:"headerPlain,omitempty"`

	SalesChannels      []SalesChannel  `json:"salesChannels,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	Id      string  `json:"id,omitempty"`

	SystemDefault      bool  `json:"systemDefault,omitempty"`

	Name      string  `json:"name,omitempty"`

	Description      string  `json:"description,omitempty"`

	FooterHtml      string  `json:"footerHtml,omitempty"`

	FooterPlain      string  `json:"footerPlain,omitempty"`

	Translations      []MailHeaderFooterTranslation  `json:"translations,omitempty"`

}

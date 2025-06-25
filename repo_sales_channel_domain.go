package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type SalesChannelDomainRepository struct {
	*GenericRepository[SalesChannelDomain]
}

func NewSalesChannelDomainRepository(client *Client) *SalesChannelDomainRepository {
	return &SalesChannelDomainRepository{
		GenericRepository: NewGenericRepository[SalesChannelDomain](client),
	}
}

func (t *SalesChannelDomainRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[SalesChannelDomain], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "sales-channel-domain")
}

func (t *SalesChannelDomainRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[SalesChannelDomain], *http.Response, error) {
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

func (t *SalesChannelDomainRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "sales-channel-domain")
}

func (t *SalesChannelDomainRepository) Upsert(ctx ApiContext, entity []SalesChannelDomain) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "sales_channel_domain")
}

func (t *SalesChannelDomainRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "sales_channel_domain")
}

type SalesChannelDomain struct {

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Currency      *Currency  `json:"currency,omitempty"`

	CurrencyId      string  `json:"currencyId,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	HreflangUseOnlyLocale      bool  `json:"hreflangUseOnlyLocale,omitempty"`

	Id      string  `json:"id,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	ProductExports      []ProductExport  `json:"productExports,omitempty"`

	SalesChannel      *SalesChannel  `json:"salesChannel,omitempty"`

	SalesChannelDefaultHreflang      *SalesChannel  `json:"salesChannelDefaultHreflang,omitempty"`

	SalesChannelId      string  `json:"salesChannelId,omitempty"`

	SnippetSet      *SnippetSet  `json:"snippetSet,omitempty"`

	SnippetSetId      string  `json:"snippetSetId,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Url      string  `json:"url,omitempty"`

}

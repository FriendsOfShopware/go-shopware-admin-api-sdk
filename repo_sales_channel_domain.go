package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type SalesChannelDomainRepository ClientService

func (t SalesChannelDomainRepository) Search(ctx ApiContext, criteria Criteria) (*SalesChannelDomainCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/sales-channel-domain", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(SalesChannelDomainCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t SalesChannelDomainRepository) SearchAll(ctx ApiContext, criteria Criteria) (*SalesChannelDomainCollection, *http.Response, error) {
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

func (t SalesChannelDomainRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/sales-channel-domain", criteria)

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

func (t SalesChannelDomainRepository) Upsert(ctx ApiContext, entity []SalesChannelDomain) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"sales_channel_domain": {
		Entity:  "sales_channel_domain",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t SalesChannelDomainRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"sales_channel_domain": {
		Entity:  "sales_channel_domain",
		Action:  "delete",
		Payload: payload,
	}})
}

type SalesChannelDomain struct {

	SalesChannelId      string  `json:"salesChannelId,omitempty"`

	CurrencyId      string  `json:"currencyId,omitempty"`

	Currency      *Currency  `json:"currency,omitempty"`

	HreflangUseOnlyLocale      bool  `json:"hreflangUseOnlyLocale,omitempty"`

	SalesChannel      *SalesChannel  `json:"salesChannel,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	SalesChannelDefaultHreflang      *SalesChannel  `json:"salesChannelDefaultHreflang,omitempty"`

	Id      string  `json:"id,omitempty"`

	SnippetSetId      string  `json:"snippetSetId,omitempty"`

	ProductExports      []ProductExport  `json:"productExports,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Url      string  `json:"url,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	SnippetSet      *SnippetSet  `json:"snippetSet,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}

type SalesChannelDomainCollection struct {
	EntityCollection

	Data []SalesChannelDomain `json:"data"`
}

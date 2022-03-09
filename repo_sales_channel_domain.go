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
	LanguageId string `json:"languageId,omitempty"`

	Currency *Currency `json:"currency,omitempty"`

	Id string `json:"id,omitempty"`

	Url string `json:"url,omitempty"`

	Language *Language `json:"language,omitempty"`

	SalesChannelDefaultHreflang *SalesChannel `json:"salesChannelDefaultHreflang,omitempty"`

	ProductExports []ProductExport `json:"productExports,omitempty"`

	HreflangUseOnlyLocale bool `json:"hreflangUseOnlyLocale,omitempty"`

	CurrencyId string `json:"currencyId,omitempty"`

	SalesChannel *SalesChannel `json:"salesChannel,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	SalesChannelId string `json:"salesChannelId,omitempty"`

	SnippetSetId string `json:"snippetSetId,omitempty"`

	SnippetSet *SnippetSet `json:"snippetSet,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`
}

type SalesChannelDomainCollection struct {
	EntityCollection

	Data []SalesChannelDomain `json:"data"`
}

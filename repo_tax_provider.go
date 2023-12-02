package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type TaxProviderRepository ClientService

func (t TaxProviderRepository) Search(ctx ApiContext, criteria Criteria) (*TaxProviderCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/tax-provider", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(TaxProviderCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t TaxProviderRepository) SearchAll(ctx ApiContext, criteria Criteria) (*TaxProviderCollection, *http.Response, error) {
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

func (t TaxProviderRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/tax-provider", criteria)

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

func (t TaxProviderRepository) Upsert(ctx ApiContext, entity []TaxProvider) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"tax_provider": {
		Entity:  "tax_provider",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t TaxProviderRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"tax_provider": {
		Entity:  "tax_provider",
		Action:  "delete",
		Payload: payload,
	}})
}

type TaxProvider struct {

	Id      string  `json:"id,omitempty"`

	AppId      string  `json:"appId,omitempty"`

	App      *App  `json:"app,omitempty"`

	Priority      float64  `json:"priority,omitempty"`

	AvailabilityRuleId      string  `json:"availabilityRuleId,omitempty"`

	AvailabilityRule      *Rule  `json:"availabilityRule,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	Active      bool  `json:"active,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Identifier      string  `json:"identifier,omitempty"`

	Name      string  `json:"name,omitempty"`

	ProcessUrl      string  `json:"processUrl,omitempty"`

	Translations      []TaxProviderTranslation  `json:"translations,omitempty"`

}

type TaxProviderCollection struct {
	EntityCollection

	Data []TaxProvider `json:"data"`
}

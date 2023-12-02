package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type TaxProviderTranslationRepository ClientService

func (t TaxProviderTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*TaxProviderTranslationCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/tax-provider-translation", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(TaxProviderTranslationCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t TaxProviderTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*TaxProviderTranslationCollection, *http.Response, error) {
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

func (t TaxProviderTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/tax-provider-translation", criteria)

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

func (t TaxProviderTranslationRepository) Upsert(ctx ApiContext, entity []TaxProviderTranslation) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"tax_provider_translation": {
		Entity:  "tax_provider_translation",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t TaxProviderTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"tax_provider_translation": {
		Entity:  "tax_provider_translation",
		Action:  "delete",
		Payload: payload,
	}})
}

type TaxProviderTranslation struct {

	LanguageId      string  `json:"languageId,omitempty"`

	TaxProvider      *TaxProvider  `json:"taxProvider,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	Name      string  `json:"name,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	TaxProviderId      string  `json:"taxProviderId,omitempty"`

}

type TaxProviderTranslationCollection struct {
	EntityCollection

	Data []TaxProviderTranslation `json:"data"`
}

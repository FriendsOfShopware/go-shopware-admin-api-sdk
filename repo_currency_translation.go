package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type CurrencyTranslationRepository ClientService

func (t CurrencyTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*CurrencyTranslationCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/currency-translation", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(CurrencyTranslationCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t CurrencyTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*CurrencyTranslationCollection, *http.Response, error) {
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

func (t CurrencyTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/currency-translation", criteria)

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

func (t CurrencyTranslationRepository) Upsert(ctx ApiContext, entity []CurrencyTranslation) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"currency_translation": {
		Entity:  "currency_translation",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t CurrencyTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"currency_translation": {
		Entity:  "currency_translation",
		Action:  "delete",
		Payload: payload,
	}})
}

type CurrencyTranslation struct {
	ShortName string `json:"shortName,omitempty"`

	Name string `json:"name,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	CurrencyId string `json:"currencyId,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	Language *Language `json:"language,omitempty"`

	Currency *Currency `json:"currency,omitempty"`
}

type CurrencyTranslationCollection struct {
	EntityCollection

	Data []CurrencyTranslation `json:"data"`
}

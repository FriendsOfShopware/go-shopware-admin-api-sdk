package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type SalutationTranslationRepository ClientService

func (t SalutationTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*SalutationTranslationCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/salutation-translation", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(SalutationTranslationCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t SalutationTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*SalutationTranslationCollection, *http.Response, error) {
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

func (t SalutationTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/salutation-translation", criteria)

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

func (t SalutationTranslationRepository) Upsert(ctx ApiContext, entity []SalutationTranslation) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"salutation_translation": {
		Entity:  "salutation_translation",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t SalutationTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"salutation_translation": {
		Entity:  "salutation_translation",
		Action:  "delete",
		Payload: payload,
	}})
}

type SalutationTranslation struct {

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Salutation      *Salutation  `json:"salutation,omitempty"`

	LetterName      string  `json:"letterName,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	DisplayName      string  `json:"displayName,omitempty"`

	SalutationId      string  `json:"salutationId,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

}

type SalutationTranslationCollection struct {
	EntityCollection

	Data []SalutationTranslation `json:"data"`
}

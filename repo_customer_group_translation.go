package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type CustomerGroupTranslationRepository ClientService

func (t CustomerGroupTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*CustomerGroupTranslationCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/customer-group-translation", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(CustomerGroupTranslationCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t CustomerGroupTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*CustomerGroupTranslationCollection, *http.Response, error) {
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

func (t CustomerGroupTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/customer-group-translation", criteria)

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

func (t CustomerGroupTranslationRepository) Upsert(ctx ApiContext, entity []CustomerGroupTranslation) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"customer_group_translation": {
		Entity:  "customer_group_translation",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t CustomerGroupTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"customer_group_translation": {
		Entity:  "customer_group_translation",
		Action:  "delete",
		Payload: payload,
	}})
}

type CustomerGroupTranslation struct {
	RegistrationIntroduction string `json:"registrationIntroduction,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	CustomerGroupId string `json:"customerGroupId,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	Name string `json:"name,omitempty"`

	RegistrationTitle string `json:"registrationTitle,omitempty"`

	RegistrationOnlyCompanyRegistration bool `json:"registrationOnlyCompanyRegistration,omitempty"`

	RegistrationSeoMetaDescription string `json:"registrationSeoMetaDescription,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	CustomerGroup *CustomerGroup `json:"customerGroup,omitempty"`

	Language *Language `json:"language,omitempty"`
}

type CustomerGroupTranslationCollection struct {
	EntityCollection

	Data []CustomerGroupTranslation `json:"data"`
}

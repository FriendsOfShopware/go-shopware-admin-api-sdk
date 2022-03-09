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
	CustomerGroupId string `json:"customerGroupId,omitempty"`

	CustomerGroup *CustomerGroup `json:"customerGroup,omitempty"`

	Language *Language `json:"language,omitempty"`

	RegistrationTitle string `json:"registrationTitle,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	RegistrationSeoMetaDescription string `json:"registrationSeoMetaDescription,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	Name string `json:"name,omitempty"`

	RegistrationIntroduction string `json:"registrationIntroduction,omitempty"`

	RegistrationOnlyCompanyRegistration bool `json:"registrationOnlyCompanyRegistration,omitempty"`
}

type CustomerGroupTranslationCollection struct {
	EntityCollection

	Data []CustomerGroupTranslation `json:"data"`
}

package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type CustomerGroupRepository ClientService

func (t CustomerGroupRepository) Search(ctx ApiContext, criteria Criteria) (*CustomerGroupCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/customer-group", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(CustomerGroupCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t CustomerGroupRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/customer-group", criteria)

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

func (t CustomerGroupRepository) Upsert(ctx ApiContext, entity []CustomerGroup) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"customer_group": {
		Entity:  "customer_group",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t CustomerGroupRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"customer_group": {
		Entity:  "customer_group",
		Action:  "delete",
		Payload: payload,
	}})
}

type CustomerGroup struct {
	Name string `json:"name,omitempty"`

	RegistrationTitle string `json:"registrationTitle,omitempty"`

	RegistrationOnlyCompanyRegistration bool `json:"registrationOnlyCompanyRegistration,omitempty"`

	DisplayGross bool `json:"displayGross,omitempty"`

	RegistrationSeoMetaDescription string `json:"registrationSeoMetaDescription,omitempty"`

	SalesChannels []SalesChannel `json:"salesChannels,omitempty"`

	Translations []CustomerGroupTranslation `json:"translations,omitempty"`

	RegistrationSalesChannels []SalesChannel `json:"registrationSalesChannels,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	RegistrationActive bool `json:"registrationActive,omitempty"`

	RegistrationIntroduction string `json:"registrationIntroduction,omitempty"`

	Customers []Customer `json:"customers,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Translated interface{} `json:"translated,omitempty"`

	Id string `json:"id,omitempty"`
}

type CustomerGroupCollection struct {
	EntityCollection

	Data []CustomerGroup `json:"data"`
}

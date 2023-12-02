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

func (t CustomerGroupRepository) SearchAll(ctx ApiContext, criteria Criteria) (*CustomerGroupCollection, *http.Response, error) {
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

	RegistrationTitle      string  `json:"registrationTitle,omitempty"`

	RegistrationSeoMetaDescription      string  `json:"registrationSeoMetaDescription,omitempty"`

	RegistrationSalesChannels      []SalesChannel  `json:"registrationSalesChannels,omitempty"`

	Name      string  `json:"name,omitempty"`

	DisplayGross      bool  `json:"displayGross,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	RegistrationOnlyCompanyRegistration      bool  `json:"registrationOnlyCompanyRegistration,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	RegistrationIntroduction      string  `json:"registrationIntroduction,omitempty"`

	Customers      []Customer  `json:"customers,omitempty"`

	SalesChannels      []SalesChannel  `json:"salesChannels,omitempty"`

	Translations      []CustomerGroupTranslation  `json:"translations,omitempty"`

	Id      string  `json:"id,omitempty"`

	RegistrationActive      bool  `json:"registrationActive,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

}

type CustomerGroupCollection struct {
	EntityCollection

	Data []CustomerGroup `json:"data"`
}

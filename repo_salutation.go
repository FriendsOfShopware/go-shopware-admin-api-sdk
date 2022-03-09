package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type SalutationRepository ClientService

func (t SalutationRepository) Search(ctx ApiContext, criteria Criteria) (*SalutationCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/salutation", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(SalutationCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t SalutationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/salutation", criteria)

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

func (t SalutationRepository) Upsert(ctx ApiContext, entity []Salutation) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"salutation": {
		Entity:  "salutation",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t SalutationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"salutation": {
		Entity:  "salutation",
		Action:  "delete",
		Payload: payload,
	}})
}

type Salutation struct {
	CustomFields interface{} `json:"customFields,omitempty"`

	SalutationKey string `json:"salutationKey,omitempty"`

	OrderCustomers []OrderCustomer `json:"orderCustomers,omitempty"`

	OrderAddresses []OrderAddress `json:"orderAddresses,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Translated interface{} `json:"translated,omitempty"`

	Id string `json:"id,omitempty"`

	Translations []SalutationTranslation `json:"translations,omitempty"`

	Customers []Customer `json:"customers,omitempty"`

	CustomerAddresses []CustomerAddress `json:"customerAddresses,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	DisplayName string `json:"displayName,omitempty"`

	LetterName string `json:"letterName,omitempty"`

	NewsletterRecipients []NewsletterRecipient `json:"newsletterRecipients,omitempty"`
}

type SalutationCollection struct {
	EntityCollection

	Data []Salutation `json:"data"`
}

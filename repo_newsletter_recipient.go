package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type NewsletterRecipientRepository ClientService

func (t NewsletterRecipientRepository) Search(ctx ApiContext, criteria Criteria) (*NewsletterRecipientCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/newsletter-recipient", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(NewsletterRecipientCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t NewsletterRecipientRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/newsletter-recipient", criteria)

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

func (t NewsletterRecipientRepository) Upsert(ctx ApiContext, entity []NewsletterRecipient) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"newsletter_recipient": {
		Entity:  "newsletter_recipient",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t NewsletterRecipientRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"newsletter_recipient": {
		Entity:  "newsletter_recipient",
		Action:  "delete",
		Payload: payload,
	}})
}

type NewsletterRecipient struct {
	SalesChannelId string `json:"salesChannelId,omitempty"`

	City string `json:"city,omitempty"`

	Hash string `json:"hash,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	Street string `json:"street,omitempty"`

	ConfirmedAt time.Time `json:"confirmedAt,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Status string `json:"status,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	SalutationId string `json:"salutationId,omitempty"`

	Salutation *Salutation `json:"salutation,omitempty"`

	SalesChannel *SalesChannel `json:"salesChannel,omitempty"`

	Title string `json:"title,omitempty"`

	FirstName string `json:"firstName,omitempty"`

	ZipCode string `json:"zipCode,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Tags []Tag `json:"tags,omitempty"`

	Language *Language `json:"language,omitempty"`

	Id string `json:"id,omitempty"`

	Email string `json:"email,omitempty"`

	LastName string `json:"lastName,omitempty"`
}

type NewsletterRecipientCollection struct {
	EntityCollection

	Data []NewsletterRecipient `json:"data"`
}
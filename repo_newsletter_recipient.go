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

func (t NewsletterRecipientRepository) SearchAll(ctx ApiContext, criteria Criteria) (*NewsletterRecipientCollection, *http.Response, error) {
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

	LanguageId      string  `json:"languageId,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Id      string  `json:"id,omitempty"`

	Title      string  `json:"title,omitempty"`

	ZipCode      string  `json:"zipCode,omitempty"`

	SalesChannel      *SalesChannel  `json:"salesChannel,omitempty"`

	Email      string  `json:"email,omitempty"`

	ConfirmedAt      time.Time  `json:"confirmedAt,omitempty"`

	Tags      []Tag  `json:"tags,omitempty"`

	Status      string  `json:"status,omitempty"`

	FirstName      string  `json:"firstName,omitempty"`

	LastName      string  `json:"lastName,omitempty"`

	Street      string  `json:"street,omitempty"`

	SalutationId      string  `json:"salutationId,omitempty"`

	Salutation      *Salutation  `json:"salutation,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	SalesChannelId      string  `json:"salesChannelId,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	City      string  `json:"city,omitempty"`

	Hash      string  `json:"hash,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

}

type NewsletterRecipientCollection struct {
	EntityCollection

	Data []NewsletterRecipient `json:"data"`
}

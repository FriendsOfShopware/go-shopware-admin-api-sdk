package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type NewsletterRecipientRepository struct {
	*GenericRepository[NewsletterRecipient]
}

func NewNewsletterRecipientRepository(client *Client) *NewsletterRecipientRepository {
	return &NewsletterRecipientRepository{
		GenericRepository: NewGenericRepository[NewsletterRecipient](client),
	}
}

func (t *NewsletterRecipientRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[NewsletterRecipient], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "newsletter-recipient")
}

func (t *NewsletterRecipientRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[NewsletterRecipient], *http.Response, error) {
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

func (t *NewsletterRecipientRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "newsletter-recipient")
}

func (t *NewsletterRecipientRepository) Upsert(ctx ApiContext, entity []NewsletterRecipient) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "newsletter_recipient")
}

func (t *NewsletterRecipientRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "newsletter_recipient")
}

type NewsletterRecipient struct {

	City      string  `json:"city,omitempty"`

	ConfirmedAt      time.Time  `json:"confirmedAt,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Email      string  `json:"email,omitempty"`

	FirstName      string  `json:"firstName,omitempty"`

	Hash      string  `json:"hash,omitempty"`

	Id      string  `json:"id,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	LastName      string  `json:"lastName,omitempty"`

	SalesChannel      *SalesChannel  `json:"salesChannel,omitempty"`

	SalesChannelId      string  `json:"salesChannelId,omitempty"`

	Salutation      *Salutation  `json:"salutation,omitempty"`

	SalutationId      string  `json:"salutationId,omitempty"`

	Status      string  `json:"status,omitempty"`

	Street      string  `json:"street,omitempty"`

	Tags      []Tag  `json:"tags,omitempty"`

	Title      string  `json:"title,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	ZipCode      string  `json:"zipCode,omitempty"`

}

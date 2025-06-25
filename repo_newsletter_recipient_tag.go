package go_shopware_admin_sdk

import (
	"net/http"

)

type NewsletterRecipientTagRepository struct {
	*GenericRepository[NewsletterRecipientTag]
}

func NewNewsletterRecipientTagRepository(client *Client) *NewsletterRecipientTagRepository {
	return &NewsletterRecipientTagRepository{
		GenericRepository: NewGenericRepository[NewsletterRecipientTag](client),
	}
}

func (t *NewsletterRecipientTagRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[NewsletterRecipientTag], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "newsletter-recipient-tag")
}

func (t *NewsletterRecipientTagRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[NewsletterRecipientTag], *http.Response, error) {
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

func (t *NewsletterRecipientTagRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "newsletter-recipient-tag")
}

func (t *NewsletterRecipientTagRepository) Upsert(ctx ApiContext, entity []NewsletterRecipientTag) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "newsletter_recipient_tag")
}

func (t *NewsletterRecipientTagRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "newsletter_recipient_tag")
}

type NewsletterRecipientTag struct {

	NewsletterRecipientId      string  `json:"newsletterRecipientId,omitempty"`

	TagId      string  `json:"tagId,omitempty"`

	NewsletterRecipient      *NewsletterRecipient  `json:"newsletterRecipient,omitempty"`

	Tag      *Tag  `json:"tag,omitempty"`

}

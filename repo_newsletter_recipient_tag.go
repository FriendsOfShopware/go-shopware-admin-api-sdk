package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type NewsletterRecipientTagRepository ClientService

func (t NewsletterRecipientTagRepository) Search(ctx ApiContext, criteria Criteria) (*NewsletterRecipientTagCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/newsletter-recipient-tag", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(NewsletterRecipientTagCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t NewsletterRecipientTagRepository) SearchAll(ctx ApiContext, criteria Criteria) (*NewsletterRecipientTagCollection, *http.Response, error) {
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

func (t NewsletterRecipientTagRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/newsletter-recipient-tag", criteria)

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

func (t NewsletterRecipientTagRepository) Upsert(ctx ApiContext, entity []NewsletterRecipientTag) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"newsletter_recipient_tag": {
		Entity:  "newsletter_recipient_tag",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t NewsletterRecipientTagRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"newsletter_recipient_tag": {
		Entity:  "newsletter_recipient_tag",
		Action:  "delete",
		Payload: payload,
	}})
}

type NewsletterRecipientTag struct {

	NewsletterRecipientId      string  `json:"newsletterRecipientId,omitempty"`

	TagId      string  `json:"tagId,omitempty"`

	NewsletterRecipient      *NewsletterRecipient  `json:"newsletterRecipient,omitempty"`

	Tag      *Tag  `json:"tag,omitempty"`

}

type NewsletterRecipientTagCollection struct {
	EntityCollection

	Data []NewsletterRecipientTag `json:"data"`
}

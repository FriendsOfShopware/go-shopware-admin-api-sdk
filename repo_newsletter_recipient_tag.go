package go_shopware_admin_sdk

import (
	"net/http"
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
	TagId string `json:"tagId,omitempty"`

	NewsletterRecipient *NewsletterRecipient `json:"newsletterRecipient,omitempty"`

	Tag *Tag `json:"tag,omitempty"`

	NewsletterRecipientId string `json:"newsletterRecipientId,omitempty"`
}

type NewsletterRecipientTagCollection struct {
	EntityCollection

	Data []NewsletterRecipientTag `json:"data"`
}

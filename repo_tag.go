package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type TagRepository ClientService

func (t TagRepository) Search(ctx ApiContext, criteria Criteria) (*TagCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/tag", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(TagCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t TagRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/tag", criteria)

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

func (t TagRepository) Upsert(ctx ApiContext, entity []Tag) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"tag": {
		Entity:  "tag",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t TagRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"tag": {
		Entity:  "tag",
		Action:  "delete",
		Payload: payload,
	}})
}

type Tag struct {
	Products []Product `json:"products,omitempty"`

	Media []Media `json:"media,omitempty"`

	Customers []Customer `json:"customers,omitempty"`

	Orders []Order `json:"orders,omitempty"`

	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	NewsletterRecipients []NewsletterRecipient `json:"newsletterRecipients,omitempty"`

	LandingPages []LandingPage `json:"landingPages,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Categories []Category `json:"categories,omitempty"`

	ShippingMethods []ShippingMethod `json:"shippingMethods,omitempty"`
}

type TagCollection struct {
	EntityCollection

	Data []Tag `json:"data"`
}

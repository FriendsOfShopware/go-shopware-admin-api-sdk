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

func (t TagRepository) SearchAll(ctx ApiContext, criteria Criteria) (*TagCollection, *http.Response, error) {
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
	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Name string `json:"name,omitempty"`

	Products []Product `json:"products,omitempty"`

	Media []Media `json:"media,omitempty"`

	ShippingMethods []ShippingMethod `json:"shippingMethods,omitempty"`

	NewsletterRecipients []NewsletterRecipient `json:"newsletterRecipients,omitempty"`

	LandingPages []LandingPage `json:"landingPages,omitempty"`

	Id string `json:"id,omitempty"`

	Categories []Category `json:"categories,omitempty"`

	Customers []Customer `json:"customers,omitempty"`

	Orders []Order `json:"orders,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`
}

type TagCollection struct {
	EntityCollection

	Data []Tag `json:"data"`
}

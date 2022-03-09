package go_shopware_admin_sdk

import (
	"net/http"
)

type LandingPageTagRepository ClientService

func (t LandingPageTagRepository) Search(ctx ApiContext, criteria Criteria) (*LandingPageTagCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/landing-page-tag", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(LandingPageTagCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t LandingPageTagRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/landing-page-tag", criteria)

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

func (t LandingPageTagRepository) Upsert(ctx ApiContext, entity []LandingPageTag) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"landing_page_tag": {
		Entity:  "landing_page_tag",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t LandingPageTagRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"landing_page_tag": {
		Entity:  "landing_page_tag",
		Action:  "delete",
		Payload: payload,
	}})
}

type LandingPageTag struct {
	LandingPageId string `json:"landingPageId,omitempty"`

	LandingPageVersionId string `json:"landingPageVersionId,omitempty"`

	TagId string `json:"tagId,omitempty"`

	LandingPage *LandingPage `json:"landingPage,omitempty"`

	Tag *Tag `json:"tag,omitempty"`
}

type LandingPageTagCollection struct {
	EntityCollection

	Data []LandingPageTag `json:"data"`
}

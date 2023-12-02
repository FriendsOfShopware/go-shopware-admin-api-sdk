package go_shopware_admin_sdk

import (
	"net/http"
	"time"
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

func (t LandingPageTagRepository) SearchAll(ctx ApiContext, criteria Criteria) (*LandingPageTagCollection, *http.Response, error) {
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

	LandingPageVersionId      string  `json:"landingPageVersionId,omitempty"`

	TagId      string  `json:"tagId,omitempty"`

	LandingPage      *LandingPage  `json:"landingPage,omitempty"`

	Tag      *Tag  `json:"tag,omitempty"`

	LandingPageId      string  `json:"landingPageId,omitempty"`

}

type LandingPageTagCollection struct {
	EntityCollection

	Data []LandingPageTag `json:"data"`
}

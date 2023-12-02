package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type LandingPageTranslationRepository ClientService

func (t LandingPageTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*LandingPageTranslationCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/landing-page-translation", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(LandingPageTranslationCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t LandingPageTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*LandingPageTranslationCollection, *http.Response, error) {
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

func (t LandingPageTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/landing-page-translation", criteria)

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

func (t LandingPageTranslationRepository) Upsert(ctx ApiContext, entity []LandingPageTranslation) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"landing_page_translation": {
		Entity:  "landing_page_translation",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t LandingPageTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"landing_page_translation": {
		Entity:  "landing_page_translation",
		Action:  "delete",
		Payload: payload,
	}})
}

type LandingPageTranslation struct {

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	LandingPageId      string  `json:"landingPageId,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	LandingPage      *LandingPage  `json:"landingPage,omitempty"`

	Name      string  `json:"name,omitempty"`

	MetaTitle      string  `json:"metaTitle,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	LandingPageVersionId      string  `json:"landingPageVersionId,omitempty"`

	Keywords      string  `json:"keywords,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	Url      string  `json:"url,omitempty"`

	SlotConfig      interface{}  `json:"slotConfig,omitempty"`

	MetaDescription      string  `json:"metaDescription,omitempty"`

}

type LandingPageTranslationCollection struct {
	EntityCollection

	Data []LandingPageTranslation `json:"data"`
}

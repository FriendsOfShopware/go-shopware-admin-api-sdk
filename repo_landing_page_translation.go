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
	SlotConfig interface{} `json:"slotConfig,omitempty"`

	MetaDescription string `json:"metaDescription,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	LandingPage *LandingPage `json:"landingPage,omitempty"`

	Name string `json:"name,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	LandingPageId string `json:"landingPageId,omitempty"`

	MetaTitle string `json:"metaTitle,omitempty"`

	Keywords string `json:"keywords,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	Language *Language `json:"language,omitempty"`

	LandingPageVersionId string `json:"landingPageVersionId,omitempty"`

	Url string `json:"url,omitempty"`
}

type LandingPageTranslationCollection struct {
	EntityCollection

	Data []LandingPageTranslation `json:"data"`
}

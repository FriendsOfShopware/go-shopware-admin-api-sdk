package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type CategoryTranslationRepository ClientService

func (t CategoryTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*CategoryTranslationCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/category-translation", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(CategoryTranslationCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t CategoryTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/category-translation", criteria)

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

func (t CategoryTranslationRepository) Upsert(ctx ApiContext, entity []CategoryTranslation) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"category_translation": {
		Entity:  "category_translation",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t CategoryTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"category_translation": {
		Entity:  "category_translation",
		Action:  "delete",
		Payload: payload,
	}})
}

type CategoryTranslation struct {
	Name string `json:"name,omitempty"`

	SlotConfig interface{} `json:"slotConfig,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	Language *Language `json:"language,omitempty"`

	CategoryVersionId string `json:"categoryVersionId,omitempty"`

	Keywords string `json:"keywords,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	CategoryId string `json:"categoryId,omitempty"`

	Category *Category `json:"category,omitempty"`

	LinkNewTab bool `json:"linkNewTab,omitempty"`

	MetaTitle string `json:"metaTitle,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Breadcrumb interface{} `json:"breadcrumb,omitempty"`

	LinkType string `json:"linkType,omitempty"`

	InternalLink string `json:"internalLink,omitempty"`

	ExternalLink string `json:"externalLink,omitempty"`

	Description string `json:"description,omitempty"`

	MetaDescription string `json:"metaDescription,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`
}

type CategoryTranslationCollection struct {
	EntityCollection

	Data []CategoryTranslation `json:"data"`
}

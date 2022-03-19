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

func (t CategoryTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*CategoryTranslationCollection, *http.Response, error) {
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
	InternalLink string `json:"internalLink,omitempty"`

	Description string `json:"description,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	CategoryVersionId string `json:"categoryVersionId,omitempty"`

	Name string `json:"name,omitempty"`

	LinkType string `json:"linkType,omitempty"`

	LinkNewTab bool `json:"linkNewTab,omitempty"`

	CategoryId string `json:"categoryId,omitempty"`

	Breadcrumb interface{} `json:"breadcrumb,omitempty"`

	SlotConfig interface{} `json:"slotConfig,omitempty"`

	ExternalLink string `json:"externalLink,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	Category *Category `json:"category,omitempty"`

	MetaTitle string `json:"metaTitle,omitempty"`

	MetaDescription string `json:"metaDescription,omitempty"`

	Keywords string `json:"keywords,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	Language *Language `json:"language,omitempty"`
}

type CategoryTranslationCollection struct {
	EntityCollection

	Data []CategoryTranslation `json:"data"`
}

package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type CmsSlotTranslationRepository ClientService

func (t CmsSlotTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*CmsSlotTranslationCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/cms-slot-translation", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(CmsSlotTranslationCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t CmsSlotTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*CmsSlotTranslationCollection, *http.Response, error) {
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

func (t CmsSlotTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/cms-slot-translation", criteria)

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

func (t CmsSlotTranslationRepository) Upsert(ctx ApiContext, entity []CmsSlotTranslation) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"cms_slot_translation": {
		Entity:  "cms_slot_translation",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t CmsSlotTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"cms_slot_translation": {
		Entity:  "cms_slot_translation",
		Action:  "delete",
		Payload: payload,
	}})
}

type CmsSlotTranslation struct {
	CustomFields interface{} `json:"customFields,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	CmsSlot *CmsSlot `json:"cmsSlot,omitempty"`

	Language *Language `json:"language,omitempty"`

	CmsSlotVersionId string `json:"cmsSlotVersionId,omitempty"`

	Config interface{} `json:"config,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	CmsSlotId string `json:"cmsSlotId,omitempty"`

	LanguageId string `json:"languageId,omitempty"`
}

type CmsSlotTranslationCollection struct {
	EntityCollection

	Data []CmsSlotTranslation `json:"data"`
}

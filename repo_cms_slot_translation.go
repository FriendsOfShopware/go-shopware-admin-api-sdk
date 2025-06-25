package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type CmsSlotTranslationRepository struct {
	*GenericRepository[CmsSlotTranslation]
}

func NewCmsSlotTranslationRepository(client *Client) *CmsSlotTranslationRepository {
	return &CmsSlotTranslationRepository{
		GenericRepository: NewGenericRepository[CmsSlotTranslation](client),
	}
}

func (t *CmsSlotTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[CmsSlotTranslation], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "cms-slot-translation")
}

func (t *CmsSlotTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[CmsSlotTranslation], *http.Response, error) {
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

func (t *CmsSlotTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "cms-slot-translation")
}

func (t *CmsSlotTranslationRepository) Upsert(ctx ApiContext, entity []CmsSlotTranslation) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "cms_slot_translation")
}

func (t *CmsSlotTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "cms_slot_translation")
}

type CmsSlotTranslation struct {

	CmsSlot      *CmsSlot  `json:"cmsSlot,omitempty"`

	CmsSlotId      string  `json:"cmsSlotId,omitempty"`

	CmsSlotVersionId      string  `json:"cmsSlotVersionId,omitempty"`

	Config      interface{}  `json:"config,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}

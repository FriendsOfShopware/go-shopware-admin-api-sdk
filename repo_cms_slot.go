package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type CmsSlotRepository struct {
	*GenericRepository[CmsSlot]
}

func NewCmsSlotRepository(client *Client) *CmsSlotRepository {
	return &CmsSlotRepository{
		GenericRepository: NewGenericRepository[CmsSlot](client),
	}
}

func (t *CmsSlotRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[CmsSlot], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "cms-slot")
}

func (t *CmsSlotRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[CmsSlot], *http.Response, error) {
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

func (t *CmsSlotRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "cms-slot")
}

func (t *CmsSlotRepository) Upsert(ctx ApiContext, entity []CmsSlot) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "cms_slot")
}

func (t *CmsSlotRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "cms_slot")
}

type CmsSlot struct {

	Block      *CmsBlock  `json:"block,omitempty"`

	BlockId      string  `json:"blockId,omitempty"`

	CmsBlockVersionId      string  `json:"cmsBlockVersionId,omitempty"`

	Config      interface{}  `json:"config,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Data      interface{}  `json:"data,omitempty"`

	FieldConfig      interface{}  `json:"fieldConfig,omitempty"`

	Id      string  `json:"id,omitempty"`

	Locked      bool  `json:"locked,omitempty"`

	Slot      string  `json:"slot,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	Translations      []CmsSlotTranslation  `json:"translations,omitempty"`

	Type      string  `json:"type,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	VersionId      string  `json:"versionId,omitempty"`

}

package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type SalesChannelTypeTranslationRepository struct {
	*GenericRepository[SalesChannelTypeTranslation]
}

func NewSalesChannelTypeTranslationRepository(client *Client) *SalesChannelTypeTranslationRepository {
	return &SalesChannelTypeTranslationRepository{
		GenericRepository: NewGenericRepository[SalesChannelTypeTranslation](client),
	}
}

func (t *SalesChannelTypeTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[SalesChannelTypeTranslation], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "sales-channel-type-translation")
}

func (t *SalesChannelTypeTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[SalesChannelTypeTranslation], *http.Response, error) {
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

func (t *SalesChannelTypeTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "sales-channel-type-translation")
}

func (t *SalesChannelTypeTranslationRepository) Upsert(ctx ApiContext, entity []SalesChannelTypeTranslation) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "sales_channel_type_translation")
}

func (t *SalesChannelTypeTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "sales_channel_type_translation")
}

type SalesChannelTypeTranslation struct {

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Description      string  `json:"description,omitempty"`

	DescriptionLong      string  `json:"descriptionLong,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	Manufacturer      string  `json:"manufacturer,omitempty"`

	Name      string  `json:"name,omitempty"`

	SalesChannelType      *SalesChannelType  `json:"salesChannelType,omitempty"`

	SalesChannelTypeId      string  `json:"salesChannelTypeId,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}

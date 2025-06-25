package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type SalesChannelTranslationRepository struct {
	*GenericRepository[SalesChannelTranslation]
}

func NewSalesChannelTranslationRepository(client *Client) *SalesChannelTranslationRepository {
	return &SalesChannelTranslationRepository{
		GenericRepository: NewGenericRepository[SalesChannelTranslation](client),
	}
}

func (t *SalesChannelTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[SalesChannelTranslation], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "sales-channel-translation")
}

func (t *SalesChannelTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[SalesChannelTranslation], *http.Response, error) {
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

func (t *SalesChannelTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "sales-channel-translation")
}

func (t *SalesChannelTranslationRepository) Upsert(ctx ApiContext, entity []SalesChannelTranslation) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "sales_channel_translation")
}

func (t *SalesChannelTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "sales_channel_translation")
}

type SalesChannelTranslation struct {

	SalesChannel      *SalesChannel  `json:"salesChannel,omitempty"`

	Language      *Language  `json:"language,omitempty"`

	Name      string  `json:"name,omitempty"`

	HomeMetaTitle      string  `json:"homeMetaTitle,omitempty"`

	HomeMetaDescription      string  `json:"homeMetaDescription,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	HomeSlotConfig      interface{}  `json:"homeSlotConfig,omitempty"`

	HomeEnabled      bool  `json:"homeEnabled,omitempty"`

	HomeName      string  `json:"homeName,omitempty"`

	HomeKeywords      string  `json:"homeKeywords,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	SalesChannelId      string  `json:"salesChannelId,omitempty"`

}

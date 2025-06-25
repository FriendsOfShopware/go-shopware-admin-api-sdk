package go_shopware_admin_sdk

import (
	"net/http"

)

type SalesChannelLanguageRepository struct {
	*GenericRepository[SalesChannelLanguage]
}

func NewSalesChannelLanguageRepository(client *Client) *SalesChannelLanguageRepository {
	return &SalesChannelLanguageRepository{
		GenericRepository: NewGenericRepository[SalesChannelLanguage](client),
	}
}

func (t *SalesChannelLanguageRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[SalesChannelLanguage], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "sales-channel-language")
}

func (t *SalesChannelLanguageRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[SalesChannelLanguage], *http.Response, error) {
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

func (t *SalesChannelLanguageRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "sales-channel-language")
}

func (t *SalesChannelLanguageRepository) Upsert(ctx ApiContext, entity []SalesChannelLanguage) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "sales_channel_language")
}

func (t *SalesChannelLanguageRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "sales_channel_language")
}

type SalesChannelLanguage struct {

	Language      *Language  `json:"language,omitempty"`

	LanguageId      string  `json:"languageId,omitempty"`

	SalesChannel      *SalesChannel  `json:"salesChannel,omitempty"`

	SalesChannelId      string  `json:"salesChannelId,omitempty"`

}

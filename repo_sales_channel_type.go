package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type SalesChannelTypeRepository struct {
	*GenericRepository[SalesChannelType]
}

func NewSalesChannelTypeRepository(client *Client) *SalesChannelTypeRepository {
	return &SalesChannelTypeRepository{
		GenericRepository: NewGenericRepository[SalesChannelType](client),
	}
}

func (t *SalesChannelTypeRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[SalesChannelType], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "sales-channel-type")
}

func (t *SalesChannelTypeRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[SalesChannelType], *http.Response, error) {
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

func (t *SalesChannelTypeRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "sales-channel-type")
}

func (t *SalesChannelTypeRepository) Upsert(ctx ApiContext, entity []SalesChannelType) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "sales_channel_type")
}

func (t *SalesChannelTypeRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "sales_channel_type")
}

type SalesChannelType struct {

	CoverUrl      string  `json:"coverUrl,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	Description      string  `json:"description,omitempty"`

	DescriptionLong      string  `json:"descriptionLong,omitempty"`

	IconName      string  `json:"iconName,omitempty"`

	Id      string  `json:"id,omitempty"`

	Manufacturer      string  `json:"manufacturer,omitempty"`

	Name      string  `json:"name,omitempty"`

	SalesChannels      []SalesChannel  `json:"salesChannels,omitempty"`

	ScreenshotUrls      interface{}  `json:"screenshotUrls,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	Translations      []SalesChannelTypeTranslation  `json:"translations,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}

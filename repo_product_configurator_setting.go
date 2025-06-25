package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type ProductConfiguratorSettingRepository struct {
	*GenericRepository[ProductConfiguratorSetting]
}

func NewProductConfiguratorSettingRepository(client *Client) *ProductConfiguratorSettingRepository {
	return &ProductConfiguratorSettingRepository{
		GenericRepository: NewGenericRepository[ProductConfiguratorSetting](client),
	}
}

func (t *ProductConfiguratorSettingRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductConfiguratorSetting], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "product-configurator-setting")
}

func (t *ProductConfiguratorSettingRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[ProductConfiguratorSetting], *http.Response, error) {
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

func (t *ProductConfiguratorSettingRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "product-configurator-setting")
}

func (t *ProductConfiguratorSettingRepository) Upsert(ctx ApiContext, entity []ProductConfiguratorSetting) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "product_configurator_setting")
}

func (t *ProductConfiguratorSettingRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "product_configurator_setting")
}

type ProductConfiguratorSetting struct {

	MediaId      string  `json:"mediaId,omitempty"`

	Position      float64  `json:"position,omitempty"`

	Media      *Media  `json:"media,omitempty"`

	Option      *PropertyGroupOption  `json:"option,omitempty"`

	CustomFields      interface{}  `json:"customFields,omitempty"`

	VersionId      string  `json:"versionId,omitempty"`

	OptionId      string  `json:"optionId,omitempty"`

	Price      interface{}  `json:"price,omitempty"`

	Product      *Product  `json:"product,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Id      string  `json:"id,omitempty"`

	ProductId      string  `json:"productId,omitempty"`

	ProductVersionId      string  `json:"productVersionId,omitempty"`

}

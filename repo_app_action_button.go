package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type AppActionButtonRepository struct {
	*GenericRepository[AppActionButton]
}

func NewAppActionButtonRepository(client *Client) *AppActionButtonRepository {
	return &AppActionButtonRepository{
		GenericRepository: NewGenericRepository[AppActionButton](client),
	}
}

func (t *AppActionButtonRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[AppActionButton], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "app-action-button")
}

func (t *AppActionButtonRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[AppActionButton], *http.Response, error) {
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

func (t *AppActionButtonRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "app-action-button")
}

func (t *AppActionButtonRepository) Upsert(ctx ApiContext, entity []AppActionButton) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "app_action_button")
}

func (t *AppActionButtonRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "app_action_button")
}

type AppActionButton struct {

	Action      string  `json:"action,omitempty"`

	App      *App  `json:"app,omitempty"`

	AppId      string  `json:"appId,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Entity      string  `json:"entity,omitempty"`

	Id      string  `json:"id,omitempty"`

	Label      string  `json:"label,omitempty"`

	Translated      interface{}  `json:"translated,omitempty"`

	Translations      []AppActionButtonTranslation  `json:"translations,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

	Url      string  `json:"url,omitempty"`

	View      string  `json:"view,omitempty"`

}

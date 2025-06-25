package go_shopware_admin_sdk

import (
	"net/http"

	"time"

)

type ScriptRepository struct {
	*GenericRepository[Script]
}

func NewScriptRepository(client *Client) *ScriptRepository {
	return &ScriptRepository{
		GenericRepository: NewGenericRepository[Script](client),
	}
}

func (t *ScriptRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[Script], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "script")
}

func (t *ScriptRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[Script], *http.Response, error) {
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

func (t *ScriptRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "script")
}

func (t *ScriptRepository) Upsert(ctx ApiContext, entity []Script) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "script")
}

func (t *ScriptRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "script")
}

type Script struct {

	Active      bool  `json:"active,omitempty"`

	App      *App  `json:"app,omitempty"`

	AppId      string  `json:"appId,omitempty"`

	CreatedAt      time.Time  `json:"createdAt,omitempty"`

	Hook      string  `json:"hook,omitempty"`

	Id      string  `json:"id,omitempty"`

	Name      string  `json:"name,omitempty"`

	Script      string  `json:"script,omitempty"`

	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`

}

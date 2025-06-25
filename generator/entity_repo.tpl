package go_shopware_admin_sdk

import (
	"net/http"
{{if .HasTimeField}}
	"time"
{{end}}
)

type {{ .FormattedName }}Repository struct {
	*GenericRepository[{{ .FormattedName }}]
}

func New{{ .FormattedName }}Repository(client *Client) *{{ .FormattedName }}Repository {
	return &{{ .FormattedName }}Repository{
		GenericRepository: NewGenericRepository[{{ .FormattedName }}](client),
	}
}

func (t *{{ .FormattedName }}Repository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[{{ .FormattedName }}], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "{{ .ApiPath }}")
}

func (t *{{ .FormattedName }}Repository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[{{ .FormattedName }}], *http.Response, error) {
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

func (t *{{ .FormattedName }}Repository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "{{ .ApiPath }}")
}

func (t *{{ .FormattedName }}Repository) Upsert(ctx ApiContext, entity []{{ .FormattedName }}) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "{{ .Name }}")
}

func (t *{{ .FormattedName }}Repository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "{{ .Name }}")
}

type {{ .FormattedName }} struct {
{{ range .Fields }}
	{{ .Key }}      {{ .Type }}  `json:"{{ .Name }},omitempty"`
{{ end }}
}

package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type {{ .FormattedName }}Repository ClientService

func (t {{ .FormattedName }}Repository) Search(ctx ApiContext, criteria Criteria) (*{{ .FormattedName }}Collection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/{{ .ApiPath }}", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new({{ .FormattedName }}Collection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t {{ .FormattedName }}Repository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/{{ .ApiPath }}", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(SearchIdsResponse)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t {{ .FormattedName }}Repository) Upsert(ctx ApiContext, entity []{{ .FormattedName }}) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"{{ .Name }}": {
		Entity:  "{{ .Name }}",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t {{ .FormattedName }}Repository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"{{ .Name }}": {
		Entity:  "{{ .Name }}",
		Action:  "delete",
		Payload: payload,
	}})
}

type {{ .FormattedName }} struct {
{{ range .Fields }}
	{{ .Key }}      {{ .Type }}  `json:"{{ .Name }},omitempty"`
{{ end }}
}

type {{ .FormattedName }}Collection struct {
	EntityCollection

	Data []{{ .FormattedName }} `json:"data"`
}

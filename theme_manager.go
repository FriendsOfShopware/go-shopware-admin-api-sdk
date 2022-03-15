package go_shopware_admin_sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
)

type ThemeManagerService ClientService

func (m ThemeManagerService) GetConfiguration(ctx ApiContext, themeId string) (*ThemeConfiguration, *http.Response, error) {
	r, err := m.Client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("/api/_action/theme/%s/configuration", themeId), nil)

	if err != nil {
		return nil, nil, errors.Wrap(err, "GetConfiguration")
	}

	var result *ThemeConfiguration

	resp, err := m.Client.Do(ctx.Context, r, &result)

	if err != nil {
		return nil, nil, err
	}

	// Old shopware version, use fields instead
	if result.CurrentFields == nil {
		result.CurrentFields = &result.Fields
	}

	return result, resp, err
}

func (m ThemeManagerService) UpdateConfiguration(ctx ApiContext, themeId string, update ThemeUpdateRequest) (*http.Response, error) {
	content, err := json.Marshal(update)

	if err != nil {
		return nil, err
	}

	r, err := m.Client.NewRequest(ctx, http.MethodPatch, fmt.Sprintf("/api/_action/theme/%s", themeId), bytes.NewReader(content))

	if err != nil {
		return nil, errors.Wrap(err, "UpdateConfiguration")
	}

	return m.Client.BareDo(ctx.Context, r)
}

type ThemeConfiguration struct {
	CurrentFields *map[string]ThemeConfigValue `json:"currentFields"`
	Fields        map[string]ThemeConfigValue  `json:"fields"`
}

type ThemeConfigValue struct {
	Value interface{} `yaml:"value" json:"value"`
}

type ThemeUpdateRequest struct {
	Config map[string]ThemeConfigValue `json:"config"`
}

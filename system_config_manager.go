package go_shopware_admin_sdk

import (
	"net/http"
	"strings"
)

type SystemConfigService ClientService

func (c SystemConfigService) UpdateConfig(ctx ApiContext, payload string) (*http.Response, error) {
	r, err := http.NewRequestWithContext(ctx.Context, http.MethodPost, c.Client.url+"/api/_action/system-config/batch", strings.NewReader(payload))

	if err != nil {
		return nil, err
	}

	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Accept", "application/json")

	return c.Client.BareDo(ctx.Context, r)
}

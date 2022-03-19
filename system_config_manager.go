package go_shopware_admin_sdk

import "net/http"

type SystemConfigService ClientService

func (c SystemConfigService) UpdateConfig(ctx ApiContext, payload interface{}) (*http.Response, error) {
	r, err := c.Client.NewRequest(ctx, http.MethodPost, "/api/_action/system-config/batch", payload)

	if err != nil {
		return nil, err
	}

	return c.Client.BareDo(ctx.Context, r)
}

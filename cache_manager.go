package go_shopware_admin_sdk

import (
	"github.com/pkg/errors"
	"net/http"
)

type CacheManagerService ClientService

func (m CacheManagerService) Clear(ctx ApiContext) (*http.Response, error) {
	r, err := m.Client.NewRequest(ctx, http.MethodDelete, "/api/_action/cache", nil)

	if err != nil {
		return nil, errors.Wrap(err, "Clear")
	}

	return m.Client.BareDo(ctx.Context, r)
}

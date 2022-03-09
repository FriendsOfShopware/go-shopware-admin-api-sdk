package go_shopware_admin_sdk

import "net/http"

type BulkService ClientService

func (b BulkService) Sync(ctx ApiContext, payload map[string]SyncOperation) (*http.Response, error) {
	req, err := b.Client.NewRequest(ctx, http.MethodPost, "/api/_action/sync", payload)

	if err != nil {
		return nil, err
	}

	return b.Client.Do(ctx.Context, req, nil)
}

type SyncOperation struct {
	Entity  string      `json:"entity"`
	Action  string      `json:"action"`
	Payload interface{} `json:"payload"`
}

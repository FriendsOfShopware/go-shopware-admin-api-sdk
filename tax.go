package go_shopware_admin_sdk

import "net/http"

type TaxRepository repository

func (t TaxRepository) Search(ctx ApiContext, criteria Criteria) (*TaxCollection, *http.Response, error) {
	req, err := t.client.NewRequest(ctx, "POST", "/api/search/tax", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(TaxCollection)
	resp, err := t.client.Do(ctx.context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

type Tax struct {
	Id      string  `json:"id"`
	Name    string  `json:"name"`
	TaxRate float32 `json:"taxRate"`
}

type TaxCollection struct {
	EntityCollection

	Data []Tax `json:"data"`
}

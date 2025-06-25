package main

import (
	"context"
	"fmt"
	"log"

	sdk "github.com/friendsofshopware/go-shopware-admin-api-sdk"
)

func main() {
	ctx := context.Background()

	creds := sdk.NewPasswordCredentials("admin", "shopware", []string{})
	client, err := sdk.NewApiClient(ctx, "https://demo.fos.gg", creds, nil)

	if err != nil {
		log.Fatalln(err)
	}

	apiContext := sdk.NewApiContext(ctx)
	criteria := sdk.Criteria{}
	criteria.Filter = []sdk.CriteriaFilter{{Type: "equals", Field: "parentId", Value: nil}}

	collection, _, _ := client.Repository.Product.Search(apiContext, criteria)

	for _, product := range collection.Data {
		fmt.Println(product.Name)
	}

	// Get current shopware version
	fmt.Println(client.Info.Info(apiContext))
}

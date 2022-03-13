package main

import (
	"context"
	"fmt"
	sdk "github.com/friendsofshopware/go-shopware-admin-api-sdk"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	creds := sdk.NewPasswordCredentials("admin", os.Getenv("ADMIN_PW"), []string{})
	client, err := sdk.NewApiClient(ctx, "https://shop.shyim.de", creds, nil)

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

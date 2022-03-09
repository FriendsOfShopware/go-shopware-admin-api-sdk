# SDK for the Shopware 6 Admin API

See example folder

## Create a client

### Password

```go
ctx := context.Background()

// Create a using password grant
creds := sdk.NewPasswordCredentials("<username>", "<password>", []string{"write"})
client, err := sdk.NewApiClient(ctx, "<url>", creds, nil)
```

### Integration

```go
ctx := context.Background()

// Create a using password grant
creds := sdk.NewIntegrationCredentials("<client-id>", "<client-secret>", []string{"write"})
client, err := sdk.NewApiClient(ctx, "<url>", creds, nil)
```

## Usage of a repository

### Search

```go
apiContext := sdk.NewApiContext(ctx)
criteria := sdk.Criteria{}

collection, _, _ := client.Repository.Tax.Search(apiContext, criteria)

for _, tax := range collection.Data {
    fmt.Println(tax.Name)
}
```

### Create/Update

```go
apiContext := sdk.NewApiContext(ctx)
client.Repository.Tax.Upsert(apiContext, []sdk.Tax{
    {TaxRate: 15, Name: "15%"},
})
```

### Delete

```go
apiContext := sdk.NewApiContext(ctx)
client.Repository.Tax.Delete(apiContext, []string{"someid"})
```
# GoLang client for Instana

Generated from the API specification at https://instana.github.io/openapi/

## Disclaimer

This is not an official product of Instana.

## Usage

```go
import "github.com/cedricziel/go-instana/instana"

// readTags will read all available application monitoring tags along with their type and category
func readTags() {
 configuration := instana.NewConfiguration()
 configuration.Host = "tenant-unit.instana.io"
 configuration.BasePath = "https://tenant-unit.instana.io"

 client := instana.NewAPIClient(configuration)
 auth := context.WithValue(context.Background(), instana.ContextAPIKey, instana.APIKey{
  Key:    apiKey,
  Prefix: "apiToken",
 })

 tags, _, err := client.ApplicationCatalogApi.GetTagsForApplication(auth)
 if err != nil {
  fmt.Fatalf("Error calling the API, aborting.")
 }

 for _, tag := range tags {
  fmt.Printf("%s (%s): %s\n", tag.Category, tag.Type, tag.Name)
 }
}
```

## License

MIT

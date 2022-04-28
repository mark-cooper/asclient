# ASClient

ASClient is an [ArchivesSpace](https://github.com/archivesspace/archivesspace) API client wrapper using [Resty](https://github.com/go-resty/resty).

## Example usage

```golang
cfg := asclient.ASpaceAPIConfig{
  URL:      "https://test.archivesspace.org/staff/api",
  Username: "admin",
  Password: "admin",
}
client := asclient.NewAPIClient(cfg)
resp, err := client.Get("repositories", map[string]string{})

if err != nil {
  t.Fatal(err.Error())
}

if resp.StatusCode() != 200 {
  t.Fatal(resp.String())
}

var repositories []Repository
json.Unmarshal([]byte(resp.String()), &repositories)

for i, repository := range repositories {
    fmt.Println(i, repository.Name)
}
```

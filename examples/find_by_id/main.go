package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/mark-cooper/asclient"
)

type APIResourcesRefResponse struct {
	Refs []struct {
		Ref string `json:"ref"`
	} `json:"resources"`
}

func main() {
	cfg := asclient.APIConfig{
		URL:      "https://test.archivesspace.org/staff/api",
		Username: "admin",
		Password: "admin",
	}
	client := asclient.NewAPIClient(cfg)
	client.Login()
	// ASpace API is confused / confusing here (value must be array formatted string)
	IDLookup := "[\"Mss002\"]"

	resp, err := client.Get("repositories/2/find_by_id/resources", asclient.QueryParams{
		Identifier: []string{IDLookup},
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var refs APIResourcesRefResponse
	json.Unmarshal([]byte(resp.String()), &refs)

	if len(refs.Refs) > 0 {
		fmt.Printf("Found ref for id %v: %v\n", IDLookup, refs.Refs[0].Ref)
	} else {
		fmt.Printf("Ref not found for id: %v\n", IDLookup)
	}
}

package main

import (
	"fmt"
	"os"

	"github.com/mark-cooper/asclient"
)

func main() {
	cfg := asclient.APIConfig{
		URL:      "https://test.archivesspace.org/staff/api",
		Username: "admin",
		Password: "admin",
	}
	client := asclient.NewAPIClient(cfg)
	client.Login()
	resp, err := client.Get("repositories/2/find_by_id/resources", asclient.QueryParams{
		// ASpace API is confused / confusing here (value must be array formatted string)
		Identifier: []string{"[\"Mss002\"]"},
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(resp.String())
}

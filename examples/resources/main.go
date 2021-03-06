package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

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

	resp, err := client.Get("repositories/2/resources", asclient.QueryParams{
		AllIds:        true,
		ModifiedSince: asclient.ModifiedSince(time.Hour * 24),
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var ids []int
	json.Unmarshal([]byte(resp.String()), &ids)

	fmt.Println(ids)
}

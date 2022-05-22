package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/mark-cooper/asclient"
)

func main() {
	srcClient, srcRepo := Setup(
		"https://test.archivesspace.org/staff/api", "admin", "admin", "YNHSC",
	)
	fmt.Printf("Source: %v\n", srcRepo)

	_, dstRepo := Setup(
		"https://test.archivesspace.org/staff/api", "admin", "admin", "PlusOne",
	)
	fmt.Printf("Destination: %v\n", dstRepo)

	ids := FetchIds(srcClient, srcRepo)

	for _, id := range ids {
		resp, err := srcClient.Get(
			fmt.Sprintf("%s/resources/%d", srcRepo.URI, id), asclient.QueryParams{},
		)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(resp.String())
	}
}

func FetchIds(client *asclient.APIClient, repo *asclient.Repository) []int {
	resp, err := client.Get(fmt.Sprintf("%s/resources", repo.URI), asclient.QueryParams{
		AllIds:        true,
		ModifiedSince: asclient.ModifiedSince(time.Hour * (24 * 7)),
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var ids []int
	json.Unmarshal([]byte(resp.String()), &ids)
	return ids
}

func Setup(url string, username string, password string, repoCode string) (*asclient.APIClient, *asclient.Repository) {
	cfg := asclient.APIConfig{
		URL:      url,
		Username: username,
		Password: password,
	}
	client := asclient.NewAPIClient(cfg)
	_, err := client.Login()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	repo, err := client.RepositoryByCode(repoCode)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return &client, &repo
}

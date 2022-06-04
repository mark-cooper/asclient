package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/mark-cooper/asclient"
)

type APIResourcesRefResponse struct {
	Refs []struct {
		Ref string `json:"ref"`
	} `json:"resources"`
}

func main() {
	srcClient, srcRepo := Setup(
		"https://test.archivesspace.org/staff/api", "admin", "admin", "YNHSC",
	)
	fmt.Printf("Source: %v\n", srcRepo)

	dstClient, dstRepo := Setup(
		"http://localhost:4567", "admin", "admin", "TEST",
	)
	fmt.Printf("Destination: %v\n", dstRepo)

	ids := FetchResourceIds(srcClient, srcRepo)

	// Process resource ids
	for _, id := range ids {
		resp, err := srcClient.Get(
			fmt.Sprintf("%s/resources/%d", srcRepo.URI, id), asclient.QueryParams{},
		)
		if err != nil {
			fmt.Println(err)
			continue
		}

		var resource asclient.Resource
		json.Unmarshal([]byte(resp.String()), &resource)

		// Move along if the resource is unpublished
		if !resource.Publish {
			continue
		}

		// Lookup the resource in destination
		identifier := srcClient.Identify(resource)
		identifierLookup := fmt.Sprintf("[\"%s\"]", identifier) // reformat for ASpace api

		resp, err = dstClient.Get(fmt.Sprintf("%s/find_by_id/resources", dstRepo.URI), asclient.QueryParams{
			Identifier: []string{identifierLookup},
		})

		if err != nil {
			fmt.Println(err)
			continue
		}

		var refs APIResourcesRefResponse
		json.Unmarshal([]byte(resp.String()), &refs)

		// If the (updated) resource already exists we need to delete it first (no overlays)
		if len(refs.Refs) > 0 {
			fmt.Printf("Deleting existing record %v: %v\n", identifierLookup, refs.Refs[0].Ref)
			_, err = dstClient.Delete(refs.Refs[0].Ref)

			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	}
}

func FetchResourceIds(client *asclient.APIClient, repo *asclient.Repository) []int {
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

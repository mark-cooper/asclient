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
	repository, err := client.RepositoryByCode("YNHSC")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(repository.RepoCode)
}

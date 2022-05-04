package asclient

import (
	"encoding/json"
	"errors"
	"path/filepath"
)

func (client *APIClient) GetRepositoryByCode(Code string) (Repository, error) {
	resp, err := client.Get("repositories", QueryParams{})

	if err != nil {
		return Repository{}, err
	}

	if resp.StatusCode() != 200 {
		return Repository{}, errors.New(string(resp.Body()))
	}

	var repository Repository
	var repositories Repositories
	json.Unmarshal([]byte(resp.String()), &repositories)

	for _, repo := range repositories {
		if repo.RepoCode == Code {
			repository = repo
		}
	}

	if repository == (Repository{}) {
		return repository, errors.New("Repository not found")
	}

	return repository, nil
}

func (client *APIClient) Login() (string, error) {
	resp, err := client.Post(
		filepath.Join("users", client.CFG.Username, "login"),
		"{}",
		QueryParams{Password: client.CFG.Password},
	)

	if err != nil {
		return "ASpaceAPIClient request error", err
	}

	if resp.StatusCode() != 200 {
		return "ASpaceAPIClient login error", errors.New(string(resp.Body()))
	}

	session := APISessionResponse{}
	json.Unmarshal(resp.Body(), &session)

	client.Headers["X-ArchivesSpace-Session"] = session.Token
	return session.Token, nil
}

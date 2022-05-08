package asclient

import (
	"encoding/json"
	"errors"
	"path/filepath"
)

func (client *APIClient) Identify(record FingerPrinter) string {
	return record.FingerPrint()
}

func (client *APIClient) Login() (string, error) {
	resp, err := client.Post(
		filepath.Join("users", client.CFG.Username, "login"),
		"{}",
		QueryParams{Password: client.CFG.Password},
	)

	if err != nil {
		return "Login error", err
	}

	session := APISessionResponse{}
	json.Unmarshal(resp.Body(), &session)

	client.Headers["X-ArchivesSpace-Session"] = session.Token
	return session.Token, nil
}

func (client *APIClient) RepositoryByCode(code string) (Repository, error) {
	resp, err := client.Get("repositories", QueryParams{})

	if err != nil {
		return Repository{}, err
	}

	var collection Collection[Repository]
	json.Unmarshal([]byte(resp.String()), &collection.Records)

	for _, record := range collection.Records {
		if client.Identify(record) == code {
			return record, nil
		}
	}

	return Repository{}, errors.New("Repository not found")
}

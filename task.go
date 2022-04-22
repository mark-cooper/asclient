package asclient

import (
	"encoding/json"
	"errors"
	"path/filepath"
)

func (client *ASpaceAPIClient) Login() (string, error) {
	resp, err := client.Post(
		filepath.Join("users", client.CFG.Username, "login"),
		"{}",
		map[string]string{
			"password": client.CFG.Password,
		},
	)

	if err != nil {
		return "ASpaceAPIClient request error", err
	}

	if resp.StatusCode() != 200 {
		return "ASpaceAPIClient login error", errors.New(string(resp.Body()))
	}

	session := ASpaceAPISessionResponse{}
	json.Unmarshal(resp.Body(), &session)

	client.Headers["X-ArchivesSpace-Session"] = session.Token
	return session.Token, nil
}

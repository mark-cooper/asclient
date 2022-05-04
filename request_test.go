package asclient

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ASpaceAPIClient_Get_Success(t *testing.T) {

	cfg := APIConfig{
		URL:      "https://test.archivesspace.org/staff/api",
		Username: "admin",
		Password: "admin",
	}
	client := NewAPIClient(cfg)
	resp, err := client.Get("repositories/2", QueryParams{})

	if err != nil {
		t.Fatal(err.Error())
	}

	if resp.StatusCode() != 200 {
		t.Fatal(resp.String())
	}

	var repository Repository
	json.Unmarshal([]byte(resp.String()), &repository)

	assert.Contains(t, resp.String(), "lock_version")
	assert.Equal(t, repository.Name, "Your Name Here Special Collection")
}

func Test_ASpaceAPIClient_CRUD(t *testing.T) {

	cfg := APIConfig{
		URL:      "https://test.archivesspace.org/staff/api",
		Username: "admin",
		Password: "admin",
	}
	client := NewAPIClient(cfg)
	client.Login()

	repository := Repository{
		RepositoryCode: "asclient_test",
		Name:           "ASCLIENT TEST",
		Publish:        false,
	}
	bytes, _ := json.Marshal(repository)
	resp, err := client.Post("repositories", string(bytes), QueryParams{})

	if err != nil {
		t.Fatal(err.Error())
	}

	if resp.StatusCode() != 200 {
		t.Fatal(resp.String())
	}

	json.Unmarshal([]byte(resp.String()), &repository)
	assert.Equal(t, "ASCLIENT TEST", repository.Name)

	repository.Name = "ASCLIENT TEST ARCHIVE"
	bytes, _ = json.Marshal(repository)
	resp, err = client.Post(repository.URI, string(bytes), QueryParams{})

	if err != nil {
		t.Fatal(err.Error())
	}

	if resp.StatusCode() != 200 {
		t.Fatal(resp.String())
	}

	json.Unmarshal([]byte(resp.String()), &repository)
	assert.Equal(t, "ASCLIENT TEST ARCHIVE", repository.Name)

	resp, err = client.Delete(repository.URI)

	if err != nil {
		t.Fatal(err.Error())
	}

	assert.Equal(t, resp.StatusCode(), 200)
}

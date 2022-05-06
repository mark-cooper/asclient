package asclient

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ASpaceAPIClient_Login_Success(t *testing.T) {

	cfg := APIConfig{
		URL:      "https://test.archivesspace.org/staff/api",
		Username: "admin",
		Password: "admin",
	}
	client := NewAPIClient(cfg)
	_, err := client.Login()

	if err != nil {
		t.Fatal(err.Error())
	}

	assert.Contains(t, client.Headers, "X-ArchivesSpace-Session")
	assert.NotEmpty(t, client.Headers["X-ArchivesSpace-Session"])
}

func Test_ASpaceAPIClient_Login_Fail_Bad_Path(t *testing.T) {

	cfg := APIConfig{
		URL:      "https://test.archivesspace.org/staff/api/xyz",
		Username: "admin",
		Password: "admin",
	}
	client := NewAPIClient(cfg)
	_, err := client.Login()

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Sinatra::NotFound")
}

func Test_ASpaceAPIClient_Login_Fail_Bad_URI(t *testing.T) {

	cfg := APIConfig{
		URL:      "xyzhttps://test.archivesspace.org/staff/api",
		Username: "admin",
		Password: "admin",
	}
	client := NewAPIClient(cfg)
	_, err := client.Login()

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unsupported protocol")
}

func Test_ASpaceAPIClient_Login_Fail_Bad_URL(t *testing.T) {

	cfg := APIConfig{
		URL:      "https://test.archivesspace.org/staff",
		Username: "admin",
		Password: "admin",
	}
	client := NewAPIClient(cfg)
	_, err := client.Login()

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Not Found")
}

func Test_ASpaceAPIClient_RepositoryByCode(t *testing.T) {

	cfg := APIConfig{
		URL:      "https://test.archivesspace.org/staff/api",
		Username: "admin",
		Password: "admin",
	}
	client := NewAPIClient(cfg)
	repository, err := client.RepositoryByCode("YNHSC")

	if err != nil {
		t.Fatal(err.Error())
	}

	assert.Equal(t, "YNHSC", repository.RepoCode)
}

package asclient

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ASpaceAPIClient_Login_Success(t *testing.T) {

	cfg := ASpaceAPIConfig{
		URL:      "https://test.archivesspace.org/staff/api",
		Username: "admin",
		Password: "admin",
	}
	client := NewAPIClient(cfg)
	_, err := client.Login()

	if err != nil {
		t.Errorf(err.Error())
	}

	assert.Contains(t, client.Headers, "X-ArchivesSpace-Session")
	assert.NotEmpty(t, client.Headers["X-ArchivesSpace-Session"])
}

func Test_ASpaceAPIClient_Login_Fail_Bad_Path(t *testing.T) {

	cfg := ASpaceAPIConfig{
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

	cfg := ASpaceAPIConfig{
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

	cfg := ASpaceAPIConfig{
		URL:      "https://test.archivesspace.org/staff",
		Username: "admin",
		Password: "admin",
	}
	client := NewAPIClient(cfg)
	_, err := client.Login()

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Not Found")
}

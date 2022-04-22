package asclient

import (
	"testing"
)

func Test_ASpaceAPIClient_Login(t *testing.T) {

	cfg := ASpaceAPIConfig{
		URL:      "https://test.archivesspace.org/staff/api",
		Username: "admin",
		Password: "admin",
	}
	client := NewAPIClient(cfg)
	client.Login()

	if client.Headers["X-ArchivesSpace-Session"] == "" {
		t.Errorf("API Session header was not set")
	}
}

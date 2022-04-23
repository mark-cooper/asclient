package asclient

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ASpaceAPIClient_Get_Success(t *testing.T) {

	cfg := ASpaceAPIConfig{
		URL:      "https://test.archivesspace.org/staff/api",
		Username: "admin",
		Password: "admin",
	}
	client := NewAPIClient(cfg)
	resp, err := client.Get("repositories", map[string]string{})

	if err != nil {
		t.Fatal(err.Error())
	}

	if resp.StatusCode() != 200 {
		t.Fatal(resp.String())
	}

	assert.Contains(t, resp.String(), "lock_version")
}

// func Test_ASpaceAPIClient_CRUD(t *testing.T) {

// 	cfg := ASpaceAPIConfig{
// 		URL:      "https://test.archivesspace.org/staff/api",
// 		Username: "admin",
// 		Password: "admin",
// 	}
// 	client := NewAPIClient(cfg)
// 	client.Login()
// }

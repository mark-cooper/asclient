package asclient

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ASpaceAPIClient_Model_Repository(t *testing.T) {
	repository := Repository{
		RepositoryCode: "test",
		Name:           "A Test Archive",
		Publish:        true,
	}
	bytes, _ := json.Marshal(repository)
	assert.Contains(t, string(bytes), repository.Name)
}

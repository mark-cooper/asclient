package asclient

import (
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
)

type APIClient struct {
	API     resty.Client
	CFG     APIConfig
	Headers map[string]string
}

type APIConfig struct {
	URL      string
	Username string
	Password string
}

type APISessionResponse struct {
	Token string `json:"session"`
}

func ModifiedSince(duration time.Duration) string {
	t := time.Now()
	timestamp := t.Add(-duration).Unix()
	return strconv.FormatInt(timestamp, 10)
}

func NewAPIClient(config APIConfig) APIClient {
	return APIClient{
		API: *resty.New(),
		CFG: config,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
}

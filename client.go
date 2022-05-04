package asclient

import (
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
)

type ASpaceAPIClient struct {
	API     resty.Client
	CFG     ASpaceAPIConfig
	Headers map[string]string
}

type ASpaceAPIConfig struct {
	URL      string
	Username string
	Password string
}

type ASpaceAPISessionResponse struct {
	Token string `json:"session"`
}

func ModifiedSince(duration time.Duration) string {
	t := time.Now()
	timestamp := t.Add(-duration).Unix()
	return strconv.FormatInt(timestamp, 10)
}

func NewAPIClient(config ASpaceAPIConfig) ASpaceAPIClient {
	return ASpaceAPIClient{
		API: *resty.New(),
		CFG: config,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
}

package asclient

import "strings"

type Collection[T any] struct {
	Records []T
}

type FingerPrinter interface {
	FingerPrint() string
}

type Repository struct {
	LockVersion int    `json:"lock_version"`
	Name        string `json:"name"`
	Publish     bool   `json:"publish,omitempty"`
	RepoCode    string `json:"repo_code"`
	URI         string `json:"uri"`
}

type Resource struct {
	LockVersion int    `json:"lock_version"`
	Title       string `json:"title"`
	Publish     bool   `json:"publish"`
	EadID       string `json:"ead_id"`
	ID0         string `json:"id_0"`
	ID1         string `json:"id_1"`
	ID2         string `json:"id_2"`
	ID3         string `json:"id_3"`
	URI         string `json:"uri"`
}

func (repository Repository) FingerPrint() string {
	return repository.RepoCode
}

func (resource Resource) FingerPrint() string {
	var parts []string
	for _, s := range []string{resource.ID0, resource.ID1, resource.ID2, resource.ID3} {
		if strings.TrimSpace(s) != "" {
			parts = append(parts, s)
		}
	}
	return strings.Join(parts, ".")
}

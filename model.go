package asclient

type FingerPrinter interface {
	FingerPrint() string
}

type Repositories []Repository

type Repository struct {
	LockVersion int    `json:"lock_version"`
	Name        string `json:"name"`
	Publish     bool   `json:"publish,omitempty"`
	RepoCode    string `json:"repo_code"`
	URI         string `json:"uri"`
}

func (repository Repository) FingerPrint() string {
	return repository.RepoCode
}

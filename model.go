package asclient

type Repository struct {
	LockVersion    int    `json:"lock_version"`
	Name           string `json:"name"`
	Publish        bool   `json:"publish,omitempty"`
	RepositoryCode string `json:"repo_code"`
	URI            string `json:"uri"`
}

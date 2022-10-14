package cmd

// Repository defines a github repo
type Repository struct {
	Name        string `json:"name"`
	Owner       string `json:"owner"`
	Description string `json:"description"`
	Private     bool   `json:"private"`
}

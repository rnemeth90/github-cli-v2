/*
Copyright © 2022 Ryan Nemeth ryannemeth<at>live<dot>com
*/

package repo

// Repository defines a github repo
type Repository struct {
	Name                      string `json:"name"`
	Owner                     string `json:"owner"`
	Description               string `json:"description"`
	Private                   bool   `json:"private"`
	Homepage                  string `json:"homepage"`
	HasIssues                 bool   `json:"has_issues"`
	HasWiki                   bool   `json:"has_wiki"`
	IsTemplate                bool   `json:"is_template"`
	DefaultBranch             bool   `json:"default_branch"`
	AllowSquashMerge          bool   `json:"allow_squash_merge"`
	AllowMergeMerge           bool   `json:"allow_merge_merge"`
	AllowRebaseMerge          bool   `json:"allow_rebase_merge"`
	AllowAutoMerge            bool   `json:"allow_auto_merge"`
	DeleteBranchOnMerge       bool   `json:"delete_branch_on_merge"`
	AllowUpdateBranch         bool   `json:"allow_update_branch"`
	UseSquashPRTitleAsDefault bool   `json:"use_squash_pr_title_as_default"`
	SquashMergeCommitTitle    string `json:"squash_merge_commit_title"`
	SquashMergeCommitMessage  string `json:"squash_merge_commit_message"`
	MergeCommitTitle          string `json:"merge_commit_title"`
	MergeCommitMessage        string `json:"merge_commit_message"`
	Archived                  bool   `json:"archived"`
	AllowForking              bool   `json:"allow_forking"`
	WebCommitSignoffRequired  bool   `json:"web_commit_signoff_required"`
}

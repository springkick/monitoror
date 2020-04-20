package models

type PullRequest struct {
	ID         int
	Title      string
	Owner      string
	Repository string
	Ref        string
}

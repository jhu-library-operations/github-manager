package api

import (
	"time"
)

type Repository struct {
	ID string
	Name string
	Description string
	URL string
	CloneURL string
	CreatedAt time.Time
	Owner RepositoryOwner

	IsPrivate bool
	HasIssuesEnabled bool
	ViewerPermission string
	DefaultBranchRef BranchRef
}

type RepositoryOwner struct {
	Login string
}

type BranchRef struct {
	Name string
}

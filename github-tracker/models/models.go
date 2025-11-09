package models

// PushEvent representa la estructura JSON completa del evento de push de GitHub.
type PushEvent struct {
	Repository Repository `json:"repository"`
	HeadCommit Commit     `json:"head_commit"` // Puede ser 'null' o un objeto
}

type Repository struct {
	FullName string `json:"full_name"`
}

// Commit representa una entrada individual en el arreglo de commits.
type Commit struct {
	ID      string     `json:"id"`
	Message string     `json:"message"`
	Author  CommitUser `json:"author"`
}

// CommitUser representa la estructura de autor o committer dentro de un commit.
type CommitUser struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

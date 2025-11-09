package entity

type Commit struct {
	ID             string `db:"id"`
	RepoName       string `db:"repo_name"`
	CommitID       string `db:"commit_id"`
	CommitMessage  string `db:"commit_message"`
	AuthorUsername string `db:"author_username"`
	AuthorEmail    string `db:"author_email"`
	Payload        string `db:"payload"`
	CreatedAt      string `db:"created_at"`
	UpdatedAt      string `db:"updated_at"`
}

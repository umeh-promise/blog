package models

type Post struct {
	ID        string   `json:"id"`
	UserID    string   `json:"user_id"`
	Title     string   `json:"title"`
	Comment   string   `json:"comment"`
	Version   int      `json:"version"`
	Tags      []string `json:"tags"`
	CratedAt  string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

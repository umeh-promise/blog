package services

import "github.com/umeh-promise/blog/internal/interfaces"

type PostService struct {
	Repo interfaces.Posts
}

func NewPostService(repo interfaces.Posts) *PostService {
	return &PostService{
		Repo: repo,
	}
}

package models

import (
	"net/http"
	"strconv"
	"strings"
)

type Post struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Version   int       `json:"-"`
	Tags      []string  `json:"tags"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	Comments  []Comment `json:"comments"`
	User      User      `json:"user"`
}

type PostWithCount struct {
	Post
	CommentsCount int `json:"comments_count"`
}

type PostPaginationQuery struct {
	Limit  int      `json:"limit" validate:"gte=1,lte=20"`
	Offset int      `json:"offset" validate:"gte=0"`
	Search string   `json:"search" validate:"max=100"`
	Sort   string   `json:"sort" validate:"oneof=asc desc"`
	Tags   []string `json:"tags" validate:"max=5"`
}

func (postQuery PostPaginationQuery) Parse(r *http.Request) (PostPaginationQuery, error) {
	queryStr := r.URL.Query()

	limit := queryStr.Get("limit")
	if limit != "" {
		lim, err := strconv.Atoi(limit)
		if err != nil {
			return postQuery, err
		}
		postQuery.Limit = lim
	}

	offset := queryStr.Get("offset")
	if offset != "" {
		off, err := strconv.Atoi(offset)
		if err != nil {
			return postQuery, err
		}
		postQuery.Offset = off
	}

	search := queryStr.Get("search")
	if search != "" {
		postQuery.Search = search
	}

	sort := queryStr.Get("sort")
	if sort != "" {
		postQuery.Sort = sort
	}

	tags := queryStr.Get("tags")
	if tags != "" {
		postQuery.Tags = strings.Split(tags, ",")
	}

	return postQuery, nil
}

package post

import (
	"errors"
	"unicode/utf8"

	"blogapi/cmd/internal"
)

var (
	ErrPostBodyEmpty        = errors.New("post body is empty")
	ErrPostBodyExceedsLimit = errors.New("post body exceeds limit")
)

type Service struct {
	Repository Repository
}

func (p Service) Create(post internal.Post) error {
	if post.Body == "" {
		return ErrPostBodyEmpty
	}

	if utf8.RuneCountInString(post.Body) > 140 {
		return ErrPostBodyExceedsLimit
	}

	return p.Repository.Insert(post)
}

func (p Service) Delete(id string) error {
	return p.Repository.Delete(id)
}

func (p Service) FindOneById(id string) (internal.Post, error) {
	return p.Repository.FindOneById(id)
}

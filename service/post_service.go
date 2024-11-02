package service

import (
	"GoSosmed/dto"
	"GoSosmed/entity"
	"GoSosmed/errorhandler"
	"GoSosmed/repository"
)

// blue print
type PostService interface {
	Create(req *dto.PostRequest) error
}

// implementasikan interface PostService
type postService struct {
	repository repository.PostRepository
}

// constructor instance postService
func NewPostService(r repository.PostRepository) *postService {
	return &postService{
		repository: r,
	}
}

// implementasi dari methode Create dalam interface PostService
func (s *postService) Create(req *dto.PostRequest) error {
	post := entity.Post{
		UserID: req.UserID,
		Tweet:  req.Tweet,
	}
	//check user mengunggah gambar
	if req.Picture != nil {
		post.PictureUrl = &req.Picture.Filename
	}

	if err := s.repository.Create(&post); err != nil {
		return &errorhandler.InternalServerError{
			Message: err.Error(),
		}
	}

	return nil
}

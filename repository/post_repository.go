package repository

import (
	"GoSosmed/entity"

	"gorm.io/gorm"
)

// blue print
type PostRepository interface {
	Create(post *entity.Post) error
}

// mengimplementasikan Interface
type postRepository struct {
	db *gorm.DB
}

// constructor postRepository
func NewPostRepository(db *gorm.DB) *postRepository {
	return &postRepository{
		db: db,
	}
}

// implementasi Create
func (r *postRepository) Create(post *entity.Post) error {
	err := r.db.Create(&post).Error

	return err
}

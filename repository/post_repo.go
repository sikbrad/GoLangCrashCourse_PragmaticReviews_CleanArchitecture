package repository

import "awesomeProject3/entity"

//cleanarch
type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}


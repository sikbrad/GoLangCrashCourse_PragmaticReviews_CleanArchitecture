package service

import (
	"awesomeProject3/entity"
	"awesomeProject3/repository"
	"errors"
	"math/rand"
)

var(
	//initialize repo with concrete firestore repository
	repo repository.PostRepository = repository.NewFirestoreRepository()
)

type PostService interface{
	Validate (post * entity.Post) error
	Create (post * entity.Post) (* entity.Post, error)
	FindAll () ([]entity.Post, error)
}

type service struct {

}


func NewPostService() PostService {
	return &service{

	}
}


func (*service) Validate(post *entity.Post) error {
	if post == nil {
		err := errors.New("The post is empty")
		return err
	}
	if post.Title == ""{
		err := errors.New("The post title is empty")
		return err
	}
	return nil
}

func (*service) Create(post *entity.Post) (*entity.Post, error) {
	post.ID = rand.Int63()
	return repo.Save(post)
}

func (*service) FindAll() ([]entity.Post, error) {
	return  repo.FindAll()
}

package service

import (
	"awesomeProject3/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

//var(
//	repo repository.PostRepository
//)

//structure will implement the postRepository interface
type MockRepository struct {
	mock.Mock
}

//implementing the post_service things
//with mockings
func (mock *MockRepository) Save(post *entity.Post) (*entity.Post, error){
	//make stubs, that returns arguments that we receive.

	args := mock.Called() // it returns arguments
	result := args.Get(0) //gets post
	// .(*entity.Post) thing is the assertion
	// and  args.Error(1) returns error, which is at idx 1.
	return result.(*entity.Post), args.Error(1)
}
func (mock *MockRepository) FindAll() ([]entity.Post, error){
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Post), args.Error(1)
}

//tesing the real function (service function)
func TestFindAll(t *testing.T){
	mockRepo := new(MockRepository)

	var identifier int64 = 1

	post := entity.Post{
		ID: identifier,
		Title: "Some Title",
		Text: "something something",
	}

	// Setup expectations
	// when "FindAll" function is called, we expect to return an array
	//
	mockRepo.On("FindAll").Return([]entity.Post{post}, nil)

	//mock repo is passed to service
	testService := NewPostService(mockRepo)
	result, _ := testService.FindAll()

	//create assertion
	// mock assertion for behavioural testing
	mockRepo.AssertExpectations(t) //why t is needed?

	// data assertion
	assert.Equal(t, post.ID, result[0].ID)
	assert.Equal(t, post.Title, result[0].Title)
	assert.Equal(t, post.Text, result[0].Text)
}


func TestValidateEmptyPost(t *testing.T){
	testingService := NewPostService(nil)

	err := testingService.Validate(nil)

	//error is expected when input is nil.
	assert.NotNil(t, err)

	assert.Equal(t, "The post is empty", err.Error())
}

func TestValidateEmptyPostTitle(t *testing.T){
	post := entity.Post{
		ID: 1,
		Title: "",
		Text: "something",
	}

	testingService := NewPostService(nil)
	err := testingService.Validate(&post)

	assert.NotNil(t, err)
	assert.Equal(t, "The post title is empty", err.Error())
}

//function name is autogen in goland ide
func TestService_Create(t *testing.T) {
	mockRepo := new(MockRepository)
	post := entity.Post{
		Title: "Sometitle11",
		Text: "What",
	}

	//setting up expectations
	mockRepo.On("Save").Return(&post, nil)

	testService := NewPostService(mockRepo)
	//var identifier int64 = 1

	result, err := testService.Create(&post)

	//add the assertion to the mock
	mockRepo.AssertExpectations(t)

	//assert.Equal(t, identifier, result.ID) -> as id is autogen random function...
	assert.NotNil(t, result.ID)
	assert.Equal(t, post.ID, result.ID)
	assert.Equal(t, post.Title, result.Title)
	assert.Equal(t, post.Text, result.Text)
	assert.Nil(t, err)
}
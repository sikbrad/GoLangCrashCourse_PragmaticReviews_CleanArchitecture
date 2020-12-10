package controller

import (
	"awesomeProject3/entity"
	"awesomeProject3/service"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"awesomeProject3/errors"
)

var (
	postService service.PostService = service.NewPostService()
)

type PostController interface {
	GetPosts (resp http.ResponseWriter, request *http.Request)
	AddPost (resp http.ResponseWriter, req *http.Request)
}

type controller struct {

}


func NewPostController() PostController{
	return &controller{}
}




func (*controller)GetPosts(resp http.ResponseWriter, request *http.Request){
	resp.Header().Set("Content-type", "application/json")
	//result, err := json.Marshal(posts)
	//_ = result

	posts, err := postService.FindAll()

	if err != nil{
		resp.WriteHeader(http.StatusInternalServerError)
		//resp.Write([]byte(`{"error":"error when marshalling the posts array"}`))
		//repo
		//resp.Write([]byte(`{"error":"error getting the posts"}`))
		//clean arch
		json.NewEncoder(resp).Encode(errors.ServiceError{
			Message : "Error when parsing message",
		})

		return
	}

	resp.WriteHeader(http.StatusOK)
	//resp.Write(result)

	//repo
	json.NewEncoder(resp).Encode(posts)

	//needed?
	//return posts
}

func (*controller)AddPost(resp http.ResponseWriter, req *http.Request){
	//var post Post
	//repo
	var post entity.Post
	err:=json.NewDecoder(req.Body).Decode(&post)

	//error when cannot parse body as post
	if err != nil{
		resp.WriteHeader(http.StatusInternalServerError)
		//resp.Write([]byte(`{"error":"error when parsing the request"}`))
		json.NewEncoder(resp).Encode(errors.ServiceError{
			Message : "error when parsing the request",
		})
		return
	}

	//add post to array
	//post.ID = len(posts) + 1
	//posts = append(posts, post)
	//log.Println("post ", post.ID, " is added")

	//repo
	post.ID = rand.Int63() //int64
	//postService.Save(&post)
	//log.Println("post ", post.ID, " is added")

	//cleanarch
	err1 := postService.Validate(&post)
	if err1 != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		//resp.Write([]byte(`{"error":"error when validating post object"}`))
		json.NewEncoder(resp).Encode(errors.ServiceError{
			Message : err1.Error(),
		})
		return
	}

	result,err2 := postService.Create(&post)
	if err2 != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		//resp.Write([]byte(`{"error":"error when validating post object"}`))
		json.NewEncoder(resp).Encode(errors.ServiceError{
			Message : "error when saving the posts",
		})
		return
	}
	log.Println("post ", post.ID, " is added")

	// send added post as response
	//reuslt,_ := json.Marshal(post)
	//
	//resp.Header().Set("Content-type", "application/json")
	//resp.WriteHeader(http.StatusOK)
	//resp.Write(reuslt)

	//2nd
	//json.NewEncoder(resp).Encode(post)
	//cleanarch3
	json.NewEncoder(resp).Encode(result)
}
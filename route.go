package main

import (
	"awesomeProject3/entity"
	"awesomeProject3/repository"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
)

//type Post struct {
//	Id int
//	Title string
//	Text string
//}



//slice is dynamic array.
//var (
//	posts []Post
//)
var (
	repo repository.PostRepository = repository.NewPostRepository()
	//posts []Post
)

//called automatically??
//func init(){
//	posts = []Post {
//		Post{
//			Id: 1,
//			Title: "title 1",
//			Text: "text 1",
//		},
//		Post{
//			Id: 2,
//			Title: "title 2",
//			Text: "text 2",
//		},
//		Post{
//			Id: 3,
//			Title: "title 3",
//			Text: "text 3",
//		},
//	}
//}

func getPosts(resp http.ResponseWriter, request *http.Request){
	resp.Header().Set("Content-type", "application/json")
	//result, err := json.Marshal(posts)
	//_ = result

	posts, err := repo.FindAll()

	if err != nil{
		resp.WriteHeader(http.StatusInternalServerError)
		//resp.Write([]byte(`{"error":"error when marshalling the posts array"}`))
		resp.Write([]byte(`{"error":"error getting the posts"}`))
		return
	}

	resp.WriteHeader(http.StatusOK)
	//resp.Write(result)

	//repo
	json.NewEncoder(resp).Encode(posts)

	//needed?
	//return posts
}

func addPost(resp http.ResponseWriter, req *http.Request){
	//var post Post
	//repo
	var post entity.Post
	err:=json.NewDecoder(req.Body).Decode(&post)

	//error when cannot parse body as post
	if err != nil{
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error":"error when parsing the request"}`))
		return
	}

	//add post to array
	//post.ID = len(posts) + 1
	//posts = append(posts, post)
	//log.Println("post ", post.ID, " is added")

	//repo
	post.ID = rand.Int63() //int64
	repo.Save(&post)
	log.Println("post ", post.ID, " is added")

	// send added post as response
	//reuslt,_ := json.Marshal(post)
	//
	//resp.Header().Set("Content-type", "application/json")
	//resp.WriteHeader(http.StatusOK)
	//resp.Write(reuslt)

	json.NewEncoder(resp).Encode(post)
}
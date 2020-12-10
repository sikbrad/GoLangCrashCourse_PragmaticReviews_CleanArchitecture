package main

import (
	"encoding/json"
	"log"
	"net/http"
)

//type Post struct {
//	Id int
//	Title string
//	Text string
//}

type Post struct {
	Id		int		`json:"id"`
	Title	string	`json:"title"`
	Text	string	`json:"text"`
}

//slice is dynamic array.
var (
	posts []Post
)

func init(){
	posts = []Post {
		Post{
			Id: 1,
			Title: "title 1",
			Text: "text 1",
		},
		Post{
			Id: 2,
			Title: "title 2",
			Text: "text 2",
		},
		Post{
			Id: 3,
			Title: "title 3",
			Text: "text 3",
		},
	}


}

func getPosts(resp http.ResponseWriter, request *http.Request){
	resp.Header().Set("Content-type", "application/json")
	result, err := json.Marshal(posts)
	//_ = result

	if err != nil{
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error":"error when marshalling the posts array"}`))
		return
	}

	resp.WriteHeader(http.StatusOK)
	resp.Write(result)
}

func addPost(resp http.ResponseWriter, req *http.Request){
	var post Post
	err:=json.NewDecoder(req.Body).Decode(&post)

	//error when cannot parse body as post
	if err != nil{
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error":"error when parsing the request"}`))
		return
	}

	//add post to array
	post.Id = len(posts) + 1
	posts = append(posts, post)
	log.Println("post ", post.Id, " is added")

	// send added post as response
	reuslt,_ := json.Marshal(post)

	resp.Header().Set("Content-type", "application/json")
	resp.WriteHeader(http.StatusOK)
	resp.Write(reuslt)
}
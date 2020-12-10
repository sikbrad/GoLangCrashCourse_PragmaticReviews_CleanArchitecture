package main


import (
	"fmt"
	"log"
	"github.com/gorilla/mux"
	"net/http"
)

func main(){
	fmt.Println("entered main")

	router := mux.NewRouter()
	const port string = ":8000"
	router.HandleFunc("/", func(resp http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(resp,"entered index")
	})
	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/posts", addPost).Methods("POST")
	log.Println("server listening on port :",port)
	log.Fatalln(http.ListenAndServe(port,router))
}
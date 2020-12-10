package main


import (
	"fmt"
	"log"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main(){
	//firestore credential
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "C:\\Users\\USER\\go\\src\\awesomeProject3\\repository\\database_key\\gqpost-3477c-firebase-adminsdk-pc8dq-bf8a0df744.json")

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
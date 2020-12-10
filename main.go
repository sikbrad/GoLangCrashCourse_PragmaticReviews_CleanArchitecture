package main

import (
	"awesomeProject3/controller"
	router "awesomeProject3/http"
	"fmt"
	"net/http"
	"os"
)

//func main(){
//	//firestore credential
//	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "C:\\Users\\USER\\go\\src\\awesomeProject3\\repository\\database_key\\gqpost-3477c-firebase-adminsdk-pc8dq-bf8a0df744.json")
//
//	fmt.Println("entered main")
//
//	router := mux.NewRouter()
//	const port string = ":8000"
//
//	router.HandleFunc("/", func(resp http.ResponseWriter, request *http.Request) {
//		fmt.Fprintln(resp,"entered index")
//	})
//	router.HandleFunc("/posts", getPosts).Methods("GET")
//	router.HandleFunc("/posts", addPost).Methods("POST")
//	log.Println("server listening on port :",port)
//	log.Fatalln(http.ListenAndServe(port,router))
//}

var(
	postController controller.PostController = controller.NewPostController()
	//httpRouter router.Router = router.NewMuxRouter()
	httpRouter router.Router = router.NewChiRouter()
)
func main(){
	//firestore credential
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "C:\\Users\\USER\\go\\src\\awesomeProject3\\repository\\database_key\\gqpost-3477c-firebase-adminsdk-pc8dq-bf8a0df744.json")

	fmt.Println("entered main")

	const port string = ":8000"

	httpRouter.GET("/", func(resp http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(resp,"entered index")
	})
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)
	httpRouter.SERVE(port)
}
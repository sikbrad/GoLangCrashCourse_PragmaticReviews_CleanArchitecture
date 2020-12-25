package main

import (
	"./controller"
	router "./http"
	"./repository"
	"./service"
	"fmt"
	"net/http"
	"os"
)


var(
	postRepository repository.PostRepository = repository.NewFirestoreRepository()
	postService service.PostService = service.NewPostService(postRepository)
	postController controller.PostController = controller.NewPostController(postService)
	//httpRouter router.Router = router.NewMuxRouter()
	httpRouter router.Router = router.NewChiRouter()
)
func main(){
	fmt.Println("program started")
	//firestore credential
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "./repository/database_key/gqpost-3477c-firebase-adminsdk-pc8dq-bf8a0df744.json")

	fmt.Println("entered main")

	const port string = ":8000"

	httpRouter.GET("/", func(resp http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(resp,"entered index")
	})
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)
	httpRouter.SERVE(port)
}
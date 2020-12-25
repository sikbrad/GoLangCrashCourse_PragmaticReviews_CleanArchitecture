package repository

import (
	"../entity"
	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"context"
	"log"
)


// this struct is implementing both methods that we defined there in interface.
type repo struct{

}

//NewPostRepository
func NewFirestoreRepository() PostRepository{
	return &repo{

	}
}

const (
	projectId = "gqpost-3477c"
	collectionName = "post"
)

func (*repo) Save(post *entity.Post) (*entity.Post, error){
	ctx := context.Background()
	client, err := firestore.NewClient(ctx,projectId)

	if err != nil {
		log.Fatalf("Failed to create a firestore client : %v", err)
		return nil, err
	}

	//close the client once we get the data
	defer client.Close()

	_, _, err = client.Collection(collectionName).Add(
		ctx,
		map[string]interface{}{
			"ID":post.ID,
			"Title":post.Title,
			"Text":post.Text,
		})
	if err != nil {
		log.Fatalf("Failed to get collection : %v", err)
		return nil, err
	}

	return post,nil
}

func (*repo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx,projectId)

	if err != nil {
		log.Fatalf("Failed to create a firestore client : %v", err)
		return nil, err
	}

	//close the client once we get the data
	defer client.Close()

	done_symbol := iterator.Done

	var posts []entity.Post
	iterator := client.Collection(collectionName).Documents(ctx)
	for {
		doc,err := iterator.Next()
		if err != nil{
			if err == done_symbol{
				break
			}
			log.Fatalf("Failed to iterate the list of posts : %v", err)
			return nil, err
		}

		// .(int) thing is type assertion.
		post := entity.Post{
			ID: doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text: doc.Data()["Text"].(string),
		}

		posts = append(posts, post)
	}
	return posts, nil
}

//Go
//import (
//"fmt"
//"context"
//
//firebase "firebase.google.com/go"
//"firebase.google.com/go/auth"
//
//"google.golang.org/api/option"
//)
//
//opt := option.WithCredentialsFile("path/to/serviceAccountKey.json")
//app, err := firebase.NewApp(context.Background(), nil, opt)
//if err != nil {
//return nil, fmt.Errorf("error initializing app: %v", err)
//}




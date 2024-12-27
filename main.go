package main

import (
	"fmt"
	"net/http"

	"web1/adapters/database"
	"web1/domain/use-cases"
	"web1/routes"
)

func main() {
	port := 8080

	// initialize the database
	db, err := database.NewInMemoryDatabase()
	//db, err := database.NewSQLiteDatabase("file:db.sqlite3")

	if err != nil {
		fmt.Printf("Error opening database: %v\n", err)
		return
	}

	// initialize use-cases
	createPost := usecases.NewCreatePostUseCase(db)
	listPosts := usecases.NewListPostsUseCase(db)
	getPost := usecases.NewGetPostUseCase(db)
	createComment := usecases.NewCreateCommentUseCase(db)

	http.HandleFunc("/posts", routes.Posts(createPost, listPosts))
	http.HandleFunc("/posts/", routes.Post(getPost))
	http.HandleFunc("/comments", routes.Comments(createComment))

	fmt.Printf("Starting server on port %v\n", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}

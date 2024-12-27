package main

import (
	"fmt"
	"net/http"

	"web1/adapters/database"
	"web1/adapters/env"
	"web1/domain/use-cases"
	"web1/routes"
)

func main() {
	env := environment.NewEnv()

	// initialize the database
	db := database.InitializeDatabase(env)

	// initialize use-cases
	createPost := usecases.NewCreatePostUseCase(db)
	listPosts := usecases.NewListPostsUseCase(db)
	getPost := usecases.NewGetPostUseCase(db)
	createComment := usecases.NewCreateCommentUseCase(db)

	// create routes
	http.HandleFunc("/posts", routes.Posts(createPost, listPosts))
	http.HandleFunc("/posts/", routes.Post(getPost))
	http.HandleFunc("/comments", routes.Comments(createComment))

	port := env.GetServerPort()
	fmt.Printf("Starting server on port %v\n", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}

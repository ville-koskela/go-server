package main

import (
    "fmt"
    "net/http"

    "web1/domain/use-cases"
    "web1/adapters/database"
    "web1/routes"
)

func main() {
    port := 8080

    // initialize the database
    db := database.NewInMemoryDatabase()

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

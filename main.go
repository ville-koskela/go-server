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
	usecases := usecases.NewUseCases(db)

	// create routes
	http.HandleFunc("/posts", routes.Posts(usecases))
	http.HandleFunc("/posts/", routes.Post(usecases))
	http.HandleFunc("/comments", routes.Comments(usecases))

	port := env.GetServerPort()
	fmt.Printf("Starting server on port %v\n", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}

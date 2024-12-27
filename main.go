package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"

    "web1/adapters/database"
    "web1/domain/models"
    "web1/domain/use-cases"
)

func main() {
    port := 8080

    createPost := usecases.NewCreatePostUseCase(database.NewInMemoryDatabase()) 

    http.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) { // Corrected here
        if r.Method == http.MethodPost {
            var p models.Post
            body, err := ioutil.ReadAll(r.Body)
            if err != nil {
                http.Error(w, "Error reading request body", http.StatusInternalServerError)
                return
            }

            err = json.Unmarshal(body, &p)
            if err != nil {
                http.Error(w, "Error unmarshalling request body", http.StatusBadRequest)
                return
            }

            post, err := createPost.Execute(&p)

            if err != nil {
                http.Error(w, "Error saving post", http.StatusInternalServerError)
                return
            }

            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusCreated)
            json.NewEncoder(w).Encode(post)
        }
    })

    fmt.Printf("Starting server on port %v\n", port) // Corrected here
    http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}

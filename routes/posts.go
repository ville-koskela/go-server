package routes 

import (
    "encoding/json"
    "io/ioutil"
    "net/http"

    "web1/domain/use-cases"
    "web1/domain/models"
)

func Posts(createPost *usecases.CreatePostUseCase) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {

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
    }
}

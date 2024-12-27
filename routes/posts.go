package routes 

import (
    "encoding/json"
    "io/ioutil"
    "net/http"

    "web1/domain/use-cases"
    "web1/domain/models"
)

func Posts(
	createPost *usecases.CreatePostUseCase, 
	listPosts *usecases.ListPostsUseCase,
	) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodGet:
            handleListCommentsRequest(w, listPosts)
        case http.MethodPost:
            handleNewPostRequest(w, r, createPost)
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    }
}

func handleNewPostRequest(w http.ResponseWriter, r *http.Request, createPost *usecases.CreatePostUseCase) {
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

func handleListCommentsRequest(w http.ResponseWriter, listPosts *usecases.ListPostsUseCase) {
    posts, err := listPosts.Execute()

    if err != nil {
        http.Error(w, "Error listing posts", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(posts)
}

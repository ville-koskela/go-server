package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"web1/domain/models"
)

type PostsUseCases interface {
	CreatePost(post *models.Post) (models.Post, error)
	ListPosts() ([]models.Post, error)
}

func Posts(usecases PostsUseCases) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handleListCommentsRequest(w, usecases)
		case http.MethodPost:
			handleNewPostRequest(w, r, usecases)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func handleNewPostRequest(w http.ResponseWriter, r *http.Request, usecases PostsUseCases) {
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

	post, err := usecases.CreatePost(&p)
	if err != nil {
		http.Error(w, "Error saving post", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}

func handleListCommentsRequest(w http.ResponseWriter, usecases PostsUseCases) {
	posts, err := usecases.ListPosts()

	if err != nil {
		http.Error(w, "Error listing posts", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

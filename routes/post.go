package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"web1/domain/models"
	"web1/domain/use-cases"
)

func Post(getPost *usecases.GetPostUseCase) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(r.URL.Path[7:], 10, 64)

		if err != nil {
			http.Error(w, "Invalid post ID", http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodGet:
			handleGetPostByIdRequest(w, id, getPost)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func handleGetPostByIdRequest(w http.ResponseWriter, id int64, getPost *usecases.GetPostUseCase) {

	post, comments, err := getPost.Execute(id)
	if err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	response := struct {
		Post     models.Post      `json:"post"`
		Comments []models.Comment `json:"comments"`
	}{
		Post:     post,
		Comments: comments,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

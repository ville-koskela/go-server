package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"web1/domain/models"
)

type PostUseCases interface {
	GetPost(id int64) (models.Post, []models.Comment, error)
}

func Post(usecases PostUseCases) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(r.URL.Path[7:], 10, 64)

		if err != nil {
			http.Error(w, "Invalid post ID", http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodGet:
			handleGetPostByIdRequest(w, id, usecases)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func handleGetPostByIdRequest(w http.ResponseWriter, id int64, usecases PostUseCases) {

	post, comments, err := usecases.GetPost(id)
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

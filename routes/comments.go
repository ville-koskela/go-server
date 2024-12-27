package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"web1/domain/models"
)

type CommentsUseCase interface {
	CreateComment(comment *models.Comment) (models.Comment, error)
}

func Comments(usecases CommentsUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handleNewCommentRequest(w, r, usecases)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func handleNewCommentRequest(w http.ResponseWriter, r *http.Request, usecases CommentsUseCase) {
	var c models.Comment
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &c)
	if err != nil {
		http.Error(w, "Error unmarshalling request body", http.StatusBadRequest)
		return
	}

	comment, err := usecases.CreateComment(&c)
	if err != nil {
		http.Error(w, "Error saving comment", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(comment)
}

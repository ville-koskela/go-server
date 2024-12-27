package routes 

import (
    "encoding/json"
    "io/ioutil"
    "net/http"

    "web1/domain/use-cases"
    "web1/domain/models"
)

func Comments(createComment *usecases.CreateCommentUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
			case http.MethodPost:
				handleNewCommentRequest(w, r, createComment)
			default:
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			}
	}
}

func handleNewCommentRequest(w http.ResponseWriter, r *http.Request, createComment *usecases.CreateCommentUseCase) {
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

	comment, err := createComment.Execute(&c)
	if err != nil {
		http.Error(w, "Error saving comment", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(comment)
}

package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sekarasiewicz/nurse/backend/pkg/models"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := &models.User{ID: 1, Name: "test 2"}
	json.NewEncoder(w).Encode(data)
}

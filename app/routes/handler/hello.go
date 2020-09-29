package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jmoiron/sqlx"
	"hsaaku.com/auth-api/app/model"
)

// Hello return Hello World
func Hello(db *sqlx.DB, w http.ResponseWriter, r *http.Request) {
	message := model.Message{
		Message: "Hello World!",
	}
	json.NewEncoder(w).Encode(message)
}

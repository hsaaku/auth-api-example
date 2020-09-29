package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	"hsaaku.com/auth-api/app/model"
)

// GetUsers ...
func GetUsers(db *sqlx.DB, w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	json.NewEncoder(w).Encode(model.Message{
		Message: fmt.Sprintf("page: %s", page),
	})

	// users := []model.User{}
	// db.Select(&users, "SELECT id, name, display_name FROM users")
	// json.NewEncoder(w).Encode(users)
}

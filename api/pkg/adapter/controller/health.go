package controller

import (
	"database/sql"
	"net/http"
)

func Health(db *sql.DB) http.HandlerFunc {
	handler := func(w http.ResponseWriter, r *http.Request) {
		err := db.Ping()
		if err != nil {
			http.Error(w, "Database connection failed", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}

	return handler
}

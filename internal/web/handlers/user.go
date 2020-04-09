package handlers

import (
	"net/http"
	"webserverREST/internal/web/binders"
)

func Parse(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" || r.Method == "PUT" {
			binders.BodyParse(r)
		}
		next.ServeHTTP(w, r)
	})
}

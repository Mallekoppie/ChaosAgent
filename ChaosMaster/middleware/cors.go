package middleware

import (
	"net/http"
)

func AllowCors(inner http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// This is bad. I'll add configuration later
		// TODO: Add configuration
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "*")
		w.Header().Add("Access-Control-Allow-Headers", "*")

		inner.ServeHTTP(w, r)
	})
}

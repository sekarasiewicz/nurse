package handlers

import "net/http"

func fallBack(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
}

func apiBuilder(fn func(http.ResponseWriter, *http.Request), allowedMethods map[string]bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, ok := allowedMethods[r.Method]; !ok {
			fallBack(w, r, http.StatusMethodNotAllowed)
		} else {
			// TODO: Authorization
			w.Header().Set("Content-Type", "application/json")
			fn(w, r)
		}
	}
}

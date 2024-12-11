package handlers

import "net/http"

func RegisterRoutes(r *http.ServeMux) {

	r.HandleFunc("/users", GetUsers)

	r.Handle("/api/", http.StripPrefix("/api", r))
}

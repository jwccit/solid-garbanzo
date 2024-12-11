package main

import (
	"flyte/goapi/internal/api/handlers"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	handlers.RegisterRoutes(mux)

	mux.Handle("/", http.FileServer(http.Dir("./static")))

	server := &http.Server{
		Addr:    ":3344",
		Handler: mux,
	}

	fmt.Println("Server is running on port 3344")
	server.ListenAndServe()
}

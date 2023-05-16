package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type MyPlugin struct{}

func (p *MyPlugin) RegisterRoutes(router chi.Router) {
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from the plugin using Chi!"))
	})

	router.Get("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		userID, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		// Perform operations with the user ID
		// ...

		w.Write([]byte("User ID: " + fmt.Sprint(userID)))
	})

	log.Println("Plugin routes registered")
}

var Plugin MyPlugin

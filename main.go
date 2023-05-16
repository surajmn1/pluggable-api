package main

import (
	"log"
	"net/http"
	"os"
	"plugin"

	"github.com/go-chi/chi"
)

type PluginAPI interface {
	RegisterRoutes(router chi.Router)
}

func main() {
	// Create a new Chi router
	router := chi.NewRouter()

	// Load the plugin
	p, err := plugin.Open("./plugins/myplugin.so")
	if err != nil {
		log.Fatal(err)
	}

	// Lookup the plugin's symbol
	symbol, err := p.Lookup("Plugin")
	if err != nil {
		log.Fatal(err)
	}

	// Assert that the symbol implements PluginAPI interface
	pluginAPI, ok := symbol.(PluginAPI)
	if !ok {
		log.Fatal("Invalid plugin")
	}

	// Register routes from the plugin
	pluginAPI.RegisterRoutes(router)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

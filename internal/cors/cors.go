package cors2

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Mangement(router *mux.Router) *cors.Cors {

	corsOptions := cors.Options{
		AllowedOrigins: []string{"http://localhost:3001"}, // Adding React app's domain here
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	}

	corsHandler := cors.New(corsOptions)

	return corsHandler
}

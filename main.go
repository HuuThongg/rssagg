package main

import (
	// "fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"net/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func main() {
	godotenv.Load()
	portString := os.Getenv("PORT") 
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))
	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err",handlerErr)
	router.Mount("/v1", v1Router)
	srv := &http.Server{
		Handler: router,
		Addr: ":" + portString,
	}
	log.Printf("server starting on port %v", portString)
	err := srv.ListenAndServe()
	if err!=nil {
		log.Fatal(err)
	}
}
package main

import (
	"github.com/go-chi/chi/v5"
	"log"
	"mail-application/handlers"
	"net/http"
)

func main() {
	log.Println("Running on port 8088...")

	r := chi.NewRouter()
	r.Post("/send-email", handlers.SendMailHandler)

	err := http.ListenAndServe(":8088", r)
	if err != nil {
		log.Panic("Error while starting server.")
	}
}

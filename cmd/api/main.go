package main

import (
	"fmt"
	"net/http"

	"github.com/gabriel-q7/go-rest-api/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {

	log.SetReportCaller(true)

	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting go api service...")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Error(err)
	}
}

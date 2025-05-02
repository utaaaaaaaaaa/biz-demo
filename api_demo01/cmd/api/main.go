package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"goapi/internal/handlers"
)

func main() {

	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Listening on port 8081...")
	err := http.ListenAndServe("localhost:8081", r)
	if err != nil {
		log.Error(err)
	}
}

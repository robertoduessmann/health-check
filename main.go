package main

import (
	"health-check/config"
	"log"
	"net/http"

	"github.com/robertoduessmann/health-check/controller"

	"github.com/gorilla/mux"
)

func main() {
	health := mux.NewRouter()
	health.Path("/status").Methods(http.MethodGet).HandlerFunc(controller.HealtCheck)

	// http.ListenAndServe(":3000", health)
	if err := http.ListenAndServe(":"+config.Get().Port, health); err != nil {
		log.Fatal(err)
	}
}

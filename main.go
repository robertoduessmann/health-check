package main

import (
	"log"
	"net/http"

	"github.com/robertoduessmann/health-check/config"

	"github.com/robertoduessmann/health-check/controller"

	"github.com/gorilla/mux"
)

func main() {
	health := mux.NewRouter()
	health.Path("/status").Methods(http.MethodGet).HandlerFunc(controller.HealtCheck)

	if err := http.ListenAndServe(":"+config.Get().Port, health); err != nil {
		log.Fatal(err)
	}
}

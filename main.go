package main

import (
	"net/http"

	"github.com/robertoduessmann/health-check/controller"

	"github.com/gorilla/mux"
)

func main() {
	health := mux.NewRouter()
	health.Path("/status").Methods(http.MethodGet).HandlerFunc(controller.HealtCheck)

	http.ListenAndServe(":3000", health)
}

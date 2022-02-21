package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lab/go-proxy/config"
	"github.com/lab/go-proxy/models"
	"github.com/lab/go-proxy/service"
	"github.com/sirupsen/logrus"
	"github.com/unrolled/secure"
)

var healthCheckRoute = models.Route{
	Name:    "HealthCheck",
	Method:  "GET",
	Pattern: "/healthz",
	HandlerFunc: func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	},
}

func main() {
	secureMiddleware := secure.New(secure.Options{IsDevelopment: true})
	r := mux.NewRouter()
	r.Use(secureMiddleware.Handler)

	r.Methods(healthCheckRoute.Method).
		Path(healthCheckRoute.Pattern).
		Name(healthCheckRoute.Name).
		Handler(healthCheckRoute.HandlerFunc)

	r.Methods(service.ReverseToFlaskRoute.Method).
		Path(service.ReverseToFlaskRoute.Pattern).
		Name(service.ReverseToFlaskRoute.Name).
		Handler(service.ReverseToFlaskRoute.HandlerFunc)

	r.Methods(service.ReverseToFlaskSvc0Route.Method).
		Path(service.ReverseToFlaskSvc0Route.Pattern).
		Name(service.ReverseToFlaskSvc0Route.Name).
		Handler(service.ReverseToFlaskSvc0Route.HandlerFunc)

	r.Methods(service.ReverseToFlaskSvc1GetRoute.Method).
		Path(service.ReverseToFlaskSvc1GetRoute.Pattern).
		Name(service.ReverseToFlaskSvc1GetRoute.Name).
		Handler(service.ReverseToFlaskSvc1GetRoute.HandlerFunc)

	r.Methods("POST", "OPTIONS").
		Path(service.ReverseToFlaskSvc1PostRoute.Pattern).
		Name(service.ReverseToFlaskSvc1PostRoute.Name).
		Handler(service.ReverseToFlaskSvc1PostRoute.HandlerFunc)

	http.Handle("/", r)
	err := http.ListenAndServe(":"+config.Port, nil)

	if err != nil {
		logrus.Errorln("An error occurred starting HTTP listener at port " + config.Port)
		logrus.Errorln("Error: " + err.Error())
	}
}

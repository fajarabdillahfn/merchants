package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	cmdAuth "github.com/fajarabdillahfn/merchants/cmd/auth"
	cmdRevenue "github.com/fajarabdillahfn/merchants/cmd/revenue"
	cDB "github.com/fajarabdillahfn/merchants/common/db/mysql"
	cMiddleware "github.com/fajarabdillahfn/merchants/common/middleware"
)

const port = ":8080"

func main() {
	log.Println("========== Start App ==========")
	initAllModules()
	router := mux.NewRouter()
	initRoutes(router)
	log.Println("========== Start Server ==========")
	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Println(err)
	}
}

func initAllModules() {
	msMerchant := cDB.OpenDB()

	cmdAuth.Initialize(msMerchant)
	cmdRevenue.Initialize(msMerchant)
}

func initRoutes(r *mux.Router) {

	//this group doesn't need authentication
	noAuth := r.NewRoute().Subrouter()
	noAuth.HandleFunc("/api/v1/login", cmdAuth.HTTPDelivery.Login).Methods("POST")

	withAuth := r.NewRoute().Subrouter()
	withAuth.Use(cMiddleware.Middleware)

	withAuth.HandleFunc("/api/v1/revenue/{dataType}", cmdRevenue.HTTPDelivery.GetRevenue).Methods("GET")
}

package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	cDB "github.com/fajarabdillahfn/merchants/common/db/mysql"
	cmdAuth "github.com/fajarabdillahfn/merchants/internal/cmd/auth"
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
}

func initRoutes(r *mux.Router) {

	//this group doesn't need authentication
	noAuth := r.NewRoute().Subrouter()
	noAuth.HandleFunc("/api/v1/login", cmdAuth.HTTPDelivery.Login).Methods("POST")
}

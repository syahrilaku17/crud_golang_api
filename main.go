package main

import (
	"awesomeProject/controllers"
	"awesomeProject/database"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	LoadAppConfig()
	database.Connect(AppConfig.ConnectionString)
	//database.Migrate()

	router := mux.NewRouter().StrictSlash(true)
	RegisterProductRoutes(router)
	log.Fatal(http.ListenAndServe(":"+AppConfig.Port, router))

	log.Print("Server started on port " + AppConfig.Port)
	log.Print("Press CTRL+C to stop")
}

func RegisterProductRoutes(router *mux.Router) {
	router.HandleFunc("/api/v1/products", controllers.GetProducts).Methods("GET")
	router.HandleFunc("/api/v1/products", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/api/v1/products/{id}", controllers.GetProductById).Methods("GET")
	router.HandleFunc("/api/v1/products/{id}", controllers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/api/v1/products/{id}", controllers.DeleteProduct).Methods("DELETE")
}

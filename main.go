package main

import (
	routeHandlers "github.com/gorilla/handlers"
	"log"
	"net/http"
	"go-todo/repository"
	"go-todo/router"
)

// our main function
func main() {

	repository.InitDB()

	// cors
	allowedHeaders := routeHandlers.AllowedHeaders([]string{"X-Requested-With", "Accept", "Content-Type", "Content-Length", "Accept-Encoding", "Accept-Language", "Authorization"})
	allowedOrigins := routeHandlers.AllowedOrigins([]string{"*"})
	allowedMethods := routeHandlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})

	// router
	routers := router.NewRouter()

	log.Fatal(http.ListenAndServe(":8888", routeHandlers.CORS(allowedHeaders, allowedMethods, allowedOrigins)(routers)))
}
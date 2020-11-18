package main

import (
	"fmt"
	"github/reccrides/Migration"
	"github/reccrides/Router"
	_ "github/reccrides/docs"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/subosito/gotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

//init env
func Init() {
	gotenv.Load()
}

// @title rideCreation
// @version 1.0
// @description This is a sample service for managing booking requests
// @termsOfService http://swagger.io/terms/
// @contact.name Yuso
// @contact.email salwa@craftfoundry.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8000
// @BasePath /
func main() {
	fmt.Println("Starting the application...")
	// Init router
	Init()
	Migration.Migration()
	// Start server
	r := Router.SetupRouter()
	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	err := http.ListenAndServe(":"+port, handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With",
		"Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"https://craftfoundry-react-app.herokuapp.com", "http://localhost:3000"}))(r))
	fmt.Println(err)
	fmt.Println("Terminating the application...")
}

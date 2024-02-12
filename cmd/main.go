package main

import (
	"net/http"

	"github.com/HarrisonHemstreet/go-ws/internal/handler"
	"github.com/HarrisonHemstreet/go-ws/internal/utils"
	_ "github.com/lib/pq"
)

func main() {
	// Example usage of `registerRoute` for your routes
	utils.RegisterRoute("/", handler.HelloWorld) // Assuming HelloWorldHandler matches the expected signature
	utils.RegisterRoute("/insertuser", handler.InsertUser)
	utils.RegisterRoute("/fetchuserbyid", handler.FetchUserByID)
	utils.RegisterRoute("/deleteuserbyid", handler.DeleteUserByID)
	utils.RegisterRoute("/updateuserbyid", handler.UpdateUserByID)

	// Start the HTTP server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

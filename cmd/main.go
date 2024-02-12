package main

import (
	"net/http"

	"github.com/HarrisonHemstreet/go-ws/internal/handler"
	"github.com/HarrisonHemstreet/go-ws/internal/utils"
	_ "github.com/lib/pq"
)

func main() {
	utils.RegisterRoute("/", handler.HelloWorld)
	utils.RegisterRoute("/insertuser", handler.InsertUser)
	utils.RegisterRoute("/fetchuserbyid", handler.FetchUserByID)
	utils.RegisterRoute("/deleteuserbyid", handler.DeleteUserByID)
	utils.RegisterRoute("/updateuserbyid", handler.UpdateUserByID)

	// Start the HTTP server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

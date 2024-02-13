package main

import (
	"net/http"

	handler "github.com/HarrisonHemstreet/go-ws/internal/handler"
	userHandler "github.com/HarrisonHemstreet/go-ws/internal/handler/user"
	"github.com/HarrisonHemstreet/go-ws/internal/middleware/jwt"
)

func main() {
	protectedMux := http.NewServeMux()
	protectedMux.HandleFunc("/user", userHandler.UserRouter)
	/* EXAMPLE FOR NEW PROTECTED ROUTES:
	  // New: Register additional protected routes on the protected sub-mux
		protectedMux.HandleFunc("/user/profile", userHandler.ProfileHandler)
		protectedMux.HandleFunc("/posts", postsHandler.PostsHandler)
	*/
	jwtProtected := jwt.ValidateToken(protectedMux)
	http.Handle("/user", jwtProtected) // Note the trailing slash to cover all subpaths
	http.HandleFunc("/login", handler.Login)

	// Start the HTTP server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

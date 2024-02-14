package utils

import (
	"net/http"

	"github.com/HarrisonHemstreet/go-ws/internal/middleware/jwt"
)

func RegisterRoutes(mainMux *http.ServeMux, pathPrefix string, specificHandler http.HandlerFunc, unprotectedMethods []string) {
	// Create a sub-mux specifically for these routes
	subMux := http.NewServeMux()
	subMux.HandleFunc("/", specificHandler) // Root of the sub-path, expects paths relative to the prefix

	// Apply JWT middleware, excluding specified methods
	jwtProtected := jwt.ValidateToken(subMux, unprotectedMethods)

	// No need to strip prefix here as we're directly associating jwtProtected with the correct pathPrefix
	mainMux.Handle(pathPrefix, jwtProtected)
}

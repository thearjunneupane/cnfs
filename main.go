package main

import (
	"log"
	"net/http"
	"os"

	"github.com/thearjunneupane/cnfs/internal/router"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	staticDir := "/public/"
	router.APIRouter.
		PathPrefix(staticDir).
		Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))
	log.Println("Listening on", ":", port)
	http.ListenAndServe(":"+port, router.APIRouter)
}

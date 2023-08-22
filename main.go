package main

import (
	"log"
	"net/http"

	"github.com/thearjunneupane/cnfs/internal/router"
)

func main() {
	port := "8080"
	staticDir := "/public/"
	router.APIRouter.
		PathPrefix(staticDir).
		Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))
	log.Println("Listening on :8080...")
	http.ListenAndServe(":"+port, router.APIRouter)
}

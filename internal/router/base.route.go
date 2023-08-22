package router

import (
	"github.com/gorilla/mux"
	"github.com/thearjunneupane/cnfs/internal/handler"
)

// APIRouter is the main router
var APIRouter = mux.NewRouter()

func init() {

	APIRouter.HandleFunc("/", handler.Auth(handler.Home)).Methods("GET", "POST")
	APIRouter.HandleFunc("/guest", handler.Home).Methods("GET", "POST")

}

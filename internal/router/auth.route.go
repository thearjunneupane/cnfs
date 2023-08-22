package router

import (
	"github.com/thearjunneupane/cnfs/internal/handler"
)

func init() {

	APIRouter.HandleFunc("/login", handler.Login).Methods("GET", "POST")
	APIRouter.HandleFunc("/logout", handler.Logout).Methods("GET", "POST")
	APIRouter.HandleFunc("/signup", handler.SignUp).Methods("GET", "POST")

}

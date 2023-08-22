package router

import (
	"github.com/thearjunneupane/cnfs/internal/handler"
)

func init() {

	APIRouter.HandleFunc("/post/create", handler.Auth(handler.CreatePost)).Methods("POST")
	APIRouter.HandleFunc("/post/{id}", handler.Auth(handler.DeletePost)).Methods("DELETE")
	APIRouter.HandleFunc("/posts", handler.Auth(handler.FetchAllPosts)).Methods("GET")
	APIRouter.HandleFunc("/posts/mine", handler.Auth(handler.FetchUserPosts)).Methods("GET")
	APIRouter.HandleFunc("/post/{id}", handler.Auth(handler.UpdatePost)).Methods("PUT")

}

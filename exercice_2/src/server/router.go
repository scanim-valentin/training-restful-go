// Package 'router' defines the behavior of each enpoint accessible by the client.
package router

import (
	"service/methods"

	"github.com/gorilla/mux"
)

// Router of the API (not setup yet)
var APIRouter *mux.Router = mux.NewRouter().StrictSlash(true)

// Setup will define each endpoint accessible by the client.
// Should be called before launching API on server side.
func Setup() {
	APIRouter.HandleFunc("/Register/", methods.Register)
	APIRouter.HandleFunc("/Login/", methods.Login)
	APIRouter.HandleFunc("/GetUserList/", methods.GetUserList)
	APIRouter.HandleFunc("/GetConversation/", methods.GetConversation)
	APIRouter.HandleFunc("/SendMessage/", methods.SendMessage)
}

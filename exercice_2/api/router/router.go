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
	APIRouter.HandleFunc("/RegisterNewUser/", methods.Register)
	APIRouter.HandleFunc("/LoginUser/", methods.Login)
	APIRouter.HandleFunc("/RetrieveOnlineUserList/", methods.GetUserList)
	APIRouter.HandleFunc("/RetrieveConversation/", methods.GetConversation)
	APIRouter.HandleFunc("/AddMessageToConversation/", methods.AddMessageToConversation)
}

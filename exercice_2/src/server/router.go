// Package 'router' defines the behavior of each enpoint accessible by the client.
package server

import (
	"service/utils"

	"github.com/gorilla/mux"
)

// Router of the API (not setup yet)
var APIRouter *mux.Router = mux.NewRouter().StrictSlash(true)

// Setup will define each endpoint accessible by the client.
// Should be called before launching API on server side.
func Setup() {
	APIRouter.
		Methods("GET").
		Path("/register").
		Queries("name", "{name}").
		HandlerFunc(utils.Register)

	APIRouter.
		Methods("GET").
		Path("/login").
		Queries("id", "{id}").
		HandlerFunc(utils.Login)

	APIRouter.
		Methods("GET").
		Path("/select").
		Queries("user", "{user}").
		Queries("other", "{other}").
		HandlerFunc(utils.GetConversation)

	// Should also be implemented on client
	APIRouter.
		Methods("POST").
		Path("/send").
		HandlerFunc(utils.SendMessage)

	APIRouter.
		Methods("GET").
		Path("/logout").
		Queries("id", "{id}").
		HandlerFunc(utils.Logout)

}

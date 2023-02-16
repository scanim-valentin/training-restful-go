package router

import (
	"service/methods"

	"github.com/gorilla/mux"
)

var APIRouter *mux.Router = mux.NewRouter().StrictSlash(true)

func Setup() {
	APIRouter.HandleFunc("/RegisterNewUser/", methods.RegisterNewUser)
	APIRouter.HandleFunc("/LoginUser/", methods.LoginUser)
	APIRouter.HandleFunc("/RetrieveOnlineUserList/", methods.RetrieveOnlineUserList)
	APIRouter.HandleFunc("/RetrieveConversation/", methods.RetrieveConversation)
	APIRouter.HandleFunc("/AddMessageToConversation/", methods.AddMessageToConversation)
}

// Package server defines the behavior of each endpoint reachable by the client.
package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"service/service/blocked"
	"service/service/contacts"
	"service/service/groups"
	"service/service/groups/members"
	"service/service/messages"
	"service/service/users"
)

// APIRouter (not setup yet)
var APIRouter *mux.Router = mux.NewRouter().StrictSlash(true)

// Setup will define each endpoint accessible by the client.
// Should be called before launching API on server side.
func Setup() {
	base := "api"
	resource := "users"
	APIRouter.
		Methods(http.MethodPost).
		Path(base + "/" + resource).
		HandlerFunc(users.AddUser).
		Name("AddUser")

	APIRouter.
		Methods(http.MethodPatch).
		Path(base + "/" + resource).
		HandlerFunc(users.ChangeStatus).
		Name("ChangeStatus")

	APIRouter.
		Methods(http.MethodDelete).
		Path(base+"/"+resource).
		Queries("id", "{id}").
		HandlerFunc(users.DeleteUser).
		Name("DeleteUser")

	resource = "contacts"
	APIRouter.
		Methods(http.MethodPost).
		Path(base + "/" + resource).
		HandlerFunc(contacts.AddToContacts).
		Name("AddToContacts")

	APIRouter.
		Methods(http.MethodGet).
		Path(base+"/"+resource).
		Queries("id", "{id}").
		HandlerFunc(contacts.GetContacts).
		Name("GetContacts")

	APIRouter.
		Methods(http.MethodDelete).
		Path(base+"/"+resource).
		Queries("userid", "{userid}", "contactid", "{contactid}").
		HandlerFunc(contacts.RemoveFromContacts).
		Name("RemoveFromContacts")

	resource = "blocked"
	APIRouter.
		Methods(http.MethodPost).
		Path(base + "/" + resource).
		HandlerFunc(blocked.BlockUser).
		Name("BlockUser")

	APIRouter.
		Methods(http.MethodGet).
		Path(base+"/"+resource).
		Queries("id", "{id}").
		HandlerFunc(blocked.GetBlockedUsers).
		Name("GetBlockedUsers")

	APIRouter.
		Methods(http.MethodDelete).
		Path(base+"/"+resource).
		Queries("userid", "{userid}", "blockedid", "{blockedid}").
		HandlerFunc(blocked.UnblockUser).
		Name("UnblockUser")

	resource = "groups"
	APIRouter.
		Methods(http.MethodPost).
		Path(base + "/" + resource).
		HandlerFunc(groups.CreateGroup).
		Name("CreateGroup")

	APIRouter.
		Methods(http.MethodGet).
		Path(base+"/"+resource).
		Queries("id", "{id}").
		HandlerFunc(groups.GetUserGroups).
		Name("GetUserGroups")

	APIRouter.
		Methods(http.MethodDelete).
		Path(base+"/"+resource).
		Queries("id", "{id}").
		HandlerFunc(groups.DeleteGroup).
		Name("DeleteGroup")

	subresource := "members"
	APIRouter.
		Methods(http.MethodPost).
		Path(base + "/" + resource + "/" + subresource).
		HandlerFunc(members.AddToGroup).
		Name("AddToGroup")

	APIRouter.
		Methods(http.MethodGet).
		Path(base+"/"+resource+"/"+subresource).
		Queries("userid", "{userid}", "groupid", "{groupid}").
		HandlerFunc(members.GetUsersInGroup).
		Name("GetUsersInGroup")

	APIRouter.
		Methods(http.MethodDelete).
		Path(base+"/"+resource+"/"+subresource).
		Queries("id", "{id}").
		HandlerFunc(members.RemoveFromGroup).
		Name("RemoveFromGroup")

	resource = "messages"
	APIRouter.
		Methods(http.MethodPost).
		Path(base + "/" + resource).
		HandlerFunc(messages.SendMessage).
		Name("SendMessage")

	APIRouter.
		Methods(http.MethodGet).
		Path(base+"/"+resource).
		Queries("id", "{id}").
		HandlerFunc(messages.GetConversation).
		Name("GetConversation")

}

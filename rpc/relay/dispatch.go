// This package is contains the handler functions needed for the Relay API
package relay

import (
	"github.com/julienschmidt/httprouter"
	"github.com/pocket_network/pocket-core/rpc/shared"
	"github.com/pocket_network/pocket-core/session"
	"net/http"
)

// Define all API handlers that are under the 'dispatch' category within this file.

/*
 "DispatchOptions" handles the localhost:<relay-port>/v1/dispatch call.
 */
func DispatchOptions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	shared.WriteResponse(w, "Hello, World!")
}

/*
 "DispatchServe" handles the localhost:<relay-port>/v1/dispatch/serve call.
 */
func DispatchServe(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	dispatch := &Dispatch{}
	shared.PopulateModelFromParams(w,r,ps,dispatch)
	if session.SearchSessionList(dispatch.DevID)!=nil{
		// Session Found
	} else {
		// Session not found
		session.CreateNewSession(dispatch.DevID)
		session.SearchSessionList(dispatch.DevID)
	}
	session.PrintSessionList()
}
package pings

import (
	"net/http"

	"github.com/sudhanshuraheja/golang-sample/pkg/appcontext"
	"github.com/sudhanshuraheja/golang-sample/pkg/responses"
)

// HTTPHandler - type to handle http
type HTTPHandler func(w http.ResponseWriter, r *http.Request)

// PingHandler - structure for handling pings and related functionality
type PingHandler struct{}

// Ping - handler for pings
func (p *PingHandler) Ping(ctx *appcontext.AppContext) HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		responses.WriteJSON(w, http.StatusOK, responses.Response{Success: "pong"})
	}
}

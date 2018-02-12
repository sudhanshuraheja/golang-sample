package server

import (
	"net/http"

	"github.com/jmoiron/sqlx"

	"github.com/gorilla/mux"
	"github.com/sudhanshuraheja/golang-sample/pkg/appcontext"
	"github.com/sudhanshuraheja/golang-sample/pkg/pings"
)

// HTTPHandler - type to handle http
type HTTPHandler func(w http.ResponseWriter, r *http.Request)

// Router - to handle routing for all URLs
func Router(ctx *appcontext.AppContext, db *sqlx.DB) http.Handler {

	pingHandler := pings.PingHandler{}

	router := mux.NewRouter()
	// GET___ .../ping
	router.HandleFunc("/ping", pingHandler.Ping(ctx)).Methods(http.MethodGet)

	return router
}

// FakeHandler - to mock requests
func FakeHandler(ctx *appcontext.AppContext, db *sqlx.DB) HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"fake\":\"yes\"}"))
	}
}

package server

import (
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"

	"github.com/jeffbmartinez/delay"
	"github.com/phyber/negroni-gzip/gzip"
	"github.com/urfave/negroni"

	"github.com/sudhanshuraheja/golang-sample/pkg/appcontext"
)

// Server - structure to hold the server
type Server struct {
	ctx *appcontext.AppContext
	db  *sqlx.DB
}

// NewServer - create a new server
func NewServer(ctx *appcontext.AppContext, db *sqlx.DB) *Server {
	return &Server{
		ctx: ctx,
		db:  db,
	}
}

// Start - start a server
func (s *Server) Start() error {
	config := s.ctx.GetConfig()
	logger := s.ctx.GetLogger()

	server := negroni.New()
	server.Use(negroni.NewRecovery())
	server.Use(negroni.NewLogger())

	router := Router(s.ctx, s.db)

	if config.EnableDelayMiddleware() {
		server.Use(delay.Middleware{})
	}

	if config.EnableGzipCompression() {
		server.Use(gzip.Gzip(gzip.DefaultCompression))
	}

	if config.EnableStaticFileServer() {
		server.Use(negroni.NewStatic(http.Dir("data")))
	}

	serverURL := fmt.Sprintf(":%s", config.Port())

	server.Use(Recover())
	server.UseHandler(router)
	logger.Infoln("Starting the server at", serverURL)
	server.Run(serverURL)

	return nil
}

// Stop - stop a server
func (s *Server) Stop() error {
	// Not sure how to stop a server
	return nil
}

// Recover - recover the server after panic
func Recover() negroni.HandlerFunc {
	return negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("Recovered from panic: %+v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}()
		next(w, r)
	})
}

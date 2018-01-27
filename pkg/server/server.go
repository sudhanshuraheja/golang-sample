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

type Server struct {
	ctx *appcontext.AppContext
	db  *sqlx.DB
}

func NewServer(ctx *appcontext.AppContext, db *sqlx.DB) *Server {
	return &Server{
		ctx: ctx,
		db:  db,
	}
}

func (s *Server) Start() error {
	config := s.ctx.GetConfig()
	logger := s.ctx.GetLogger()

	server := negroni.New()
	server.Use(negroni.NewRecovery())
	server.Use(negroni.NewLogger())

	router := Router()

	if config.EnableDelayMiddleware() {
		server.Use(delay.Middleware{})
	}

	if config.EnableGzipCompression() {
		server.Use(gzip.Gzip(gzip.DefaultCompression))
	}

	if config.EnableStaticFileServer() {
		server.Use(negroni.NewStatic(http.Dir("data")))
	}

	server.Use(Recover())
	server.UseHandler(router)

	serverURL := fmt.Sprintf(":%s", config.Port())
	logger.Infoln("The server is now running at", serverURL)
	return http.ListenAndServe(serverURL, server)
}

func (s *Server) Stop() error {
	// Not sure how to stop a server
	return nil
}

func Recover() negroni.HandlerFunc {
	return negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("Recovered from panic: %+v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}()
		next(w, r)
	})
}

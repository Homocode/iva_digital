package api

import (
	"net/http"

	db "github.com/homocode/libro_iva_afip/db"
)

type Server struct {
	httpServer *http.Server
	store      *db.SQLStore
}

// Creates a new HTTP server and setup routing
func NewServer(store *db.SQLStore, address string) *Server {

	server := &Server{
		store: store,
		httpServer: &http.Server{
			Addr: address,
		},
	}

	return server
}

func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}

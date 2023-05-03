package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/homocode/libro_iva_afip/db"
)

type Server struct {
	HttpServer *http.Server
	Store      *db.SQLStore
}

func NewServer(store *db.SQLStore, addr string) *Server {
	router := mux.NewRouter()

	s := &Server{
		HttpServer: &http.Server{
			Addr:    addr,
			Handler: router,
		},
		Store: store,
	}

	s.routes(router)

	return s
}

func (s *Server) routes(r *mux.Router) {
	r.HandleFunc("/clientes", s.ListClientes()).Methods("GET")
	r.HandleFunc("/clientes", s.CreateCliente()).Methods("POST")
	r.HandleFunc("/upload", s.UploadFile()).Methods("POST")
}

/* type Server struct {
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
	http.HandleFunc("")
	return s.httpServer.ListenAndServe()
}
*/

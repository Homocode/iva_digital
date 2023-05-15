package api

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/homocode/libro_iva_afip/db"
)

type Server struct {
	HttpServer *http.Server
	Store      *db.SQLStore
}

func NewServer(store *db.SQLStore, addr string) *Server {
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:4200"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
	)
	router := mux.NewRouter()

	s := &Server{
		HttpServer: &http.Server{
			Addr:    addr,
			Handler: corsHandler(router),
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

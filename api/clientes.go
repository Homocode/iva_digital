package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/homocode/libro_iva_afip/db"
)

func (s *Server) CreateCliente() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var c db.Clientes
		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println(">>", c)
		newCliente, err := s.Store.CreateCliente(context.Background(), db.CreateClienteParams(c))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err = json.NewEncoder(w).Encode(newCliente); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) ListClientes() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		clientes, err := s.Store.ListClientes(context.Background())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err = json.NewEncoder(w).Encode(clientes); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

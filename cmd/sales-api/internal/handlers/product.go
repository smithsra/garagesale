package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	"github.com/smithsra/garagesale/internal/product"
)

// Product has handler methods for dealing with Products.
type Product struct {
	DB  *sqlx.DB
	Log *log.Logger
}

// List lists all products.
// If you open localhost:8000 in your browser, you may notice
// double requets being made. This happens because the browser
// sends a request in the background for a website favicon.
func (p *Product) List(w http.ResponseWriter, r *http.Request) {

	list, err := product.List(p.DB)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		p.Log.Println("error querying db", err)
		return
	}

	data, err := json.Marshal(list)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		p.Log.Println("error marshalling", err)
		return
	}

	w.Header().Set("content-type", "application/json: charset=utf-8")
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(data); err != nil {
		p.Log.Println("error writing", err)
	}
}

// Retrieve gives a single product.
func (p *Product) Retrieve(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	prod, err := product.Retrieve(p.DB, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		p.Log.Println("error querying db", err)
		return
	}

	data, err := json.Marshal(prod)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		p.Log.Println("error marshalling", err)
		return
	}

	w.Header().Set("content-type", "application/json: charset=utf-8")
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(data); err != nil {
		p.Log.Println("error writing", err)
	}
}

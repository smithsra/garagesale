package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	"github.com/smithsra/garagesale/internal/platform/web"
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
func (p *Product) List(w http.ResponseWriter, r *http.Request) error {

	list, err := product.List(p.DB)
	if err != nil {
		return err
	}

	return web.Respond(w, list, http.StatusOK)
}

// Retrieve gives a single product.
func (p *Product) Retrieve(w http.ResponseWriter, r *http.Request) error {

	id := chi.URLParam(r, "id")

	prod, err := product.Retrieve(p.DB, id)
	if err != nil {
		return err
	}

	return web.Respond(w, prod, http.StatusOK)
}

// Create decodes a JSON document from a POST requets and creates a new product
func (p *Product) Create(w http.ResponseWriter, r *http.Request) error {

	var np product.NewProduct

	if err := web.Decode(r, &np); err != nil {
		return err
	}

	prod, err := product.Create(p.DB, np, time.Now())
	if err != nil {
		return err
	}

	return web.Respond(w, prod, http.StatusCreated)
}

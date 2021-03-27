package web

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

// Handler is the signature that all application handlers will implement
type Handler func(http.ResponseWriter, *http.Request) error

// App is the entry point for all web applications
type App struct {
	mux *chi.Mux
	log *log.Logger
}

// NewApp knows how to contruct an internal state for an App
func NewApp(logger *log.Logger) *App {
	return &App{
		mux: chi.NewRouter(),
		log: logger,
	}
}

// Handle connects a method and a pattern to a particular application handler
func (a *App) Handle(method, pattern string, h Handler) {

	fn := func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			a.log.Printf("ERROR : %v", err)

			if err := RespondError(w, err); err != nil {
				a.log.Printf("ERROR : %v", err)
			}
		}
	}

	a.mux.MethodFunc(method, pattern, fn)
}

// ServeHTTP
func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}

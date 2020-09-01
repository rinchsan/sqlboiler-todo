package todo

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi"
)

func NewRouter(db *sql.DB) http.Handler {
	r := chi.NewRouter()
	h := newHandler(db)

	r.Get("/", h.GetAll)
	r.Post("/", h.Add)
	r.Put("/", h.Update)

	return r
}

type handler struct {
	db *sql.DB
}

func newHandler(db *sql.DB) handler {
	return handler{
		db: db,
	}
}

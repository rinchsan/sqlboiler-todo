package main

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/rinchsan/sqlboiler-todo/cmd/api/adapter/healthcheck"
	"github.com/rinchsan/sqlboiler-todo/cmd/api/adapter/todo"
	"github.com/rinchsan/sqlboiler-todo/cmd/api/adapter/user"
	"github.com/rinchsan/sqlboiler-todo/pkg/conn"
	"github.com/rinchsan/sqlboiler-todo/pkg/logger"
	"github.com/rinchsan/sqlboiler-todo/pkg/server"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func main() {
	flush, err := logger.Setup()
	if err != nil {
		panic(err)
	}
	defer flush()

	boil.DebugMode = true

	db := conn.NewDB()
	defer db.Close()

	router := newRouter(db)

	code := server.Run(router, 8080)
	os.Exit(code)
}

func newRouter(db *sql.DB) http.Handler {
	r := chi.NewRouter()

	r.Mount("/", healthcheck.NewRouter())
	r.Mount("/users", user.NewRouter(db))
	r.Mount("/todos", todo.NewRouter(db))

	return r
}

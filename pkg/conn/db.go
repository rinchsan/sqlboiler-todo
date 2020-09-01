package conn

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rinchsan/sqlboiler-todo/pkg/config"
)

func NewDB() *sql.DB {
	db, err := sql.Open(config.DB.Driver, config.DB.DSN)
	if err != nil {
		panic(err)
	}

	maxOpenConns := 30
	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxOpenConns)
	db.SetConnMaxLifetime(time.Duration(maxOpenConns) * time.Second)

	return db
}

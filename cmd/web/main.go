package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"
	"snippetbox-m0ta-b1lla/internal/models"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	logger   *slog.Logger
	snippets *models.SnippetModel
}

func main() {

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	dsn := flag.String("dsn", "web:2401@/snippetbox?parseTime=true", "entry point of sql database")
	addr := flag.String("addr", ":4000", "http port value")
	flag.Parse()
	db, err := OpenDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	app := &application{
		logger:   logger,
		snippets: &models.SnippetModel{DB: db},
	}

	logger.Info("starting server", "addr", *addr)

	defer db.Close()
	err = http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
func OpenDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}

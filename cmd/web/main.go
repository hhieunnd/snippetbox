package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/hhieunnd/snippetbox/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type config struct {
	addr      string
	staticDir string
}

type application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
	snippets *models.SnippetModel
}

func main() {
	var cfg config
	flag.StringVar(&cfg.addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&cfg.staticDir, "static-dir", "./ui/static", "Path to static assets")
	dsn := flag.String("dsn", "postgres://postgres:It123456@@localhost:5432/postgres", "PostgreSQL DSN")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	app := &application{
		infoLog:  infoLog,
		errorLog: errorLog,
		snippets: &models.SnippetModel{DB: db},
	}

	srv := http.Server{
		Addr:     cfg.addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", cfg.addr)

	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*pgxpool.Pool, error) {
	// Connect using pool
	db, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	// Check connect
	err = db.Ping(context.Background())
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

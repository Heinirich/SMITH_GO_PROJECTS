package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

const VERSION = "1.0.0"

const CSSVERSION = "1"

type config struct {
	port int
	env  string
	api  string
	db   struct {
		dsn string
	}
	stripe struct {
		secret string
		key    string
	}
}

type application struct {
	config  config
	infoLog *log.Logger
	errorLog *log.Logger
	templateCache map[string] *template.Template
	version string
}

func (app *application) serve() error{

	srv:= &http.Server{
		Addr: fmt.Sprintf(":%d", app.config.port),
		Handler: app.routes(),
		IdleTimeout: 30 * time.Second,
		ReadTimeout: 10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	app.infoLog.Printf("Starting server on port %d", app.config.port)

	return srv.ListenAndServe()

}

func main() {
	var cfg config

	// Port is the TCP port the server listens on
	flag.IntVar(&cfg.port, "port", 4000, "Server port to listen on")

	// Env is the application environment (development or production)
	flag.StringVar(&cfg.env, "env", "development", "Application environment (development|production)")

	// API is the URL to the API
	flag.StringVar(&cfg.api,"api","http://localhost:40001","URL to API")

	flag.Parse()

	cfg.stripe.key = os.Getenv("STRIPE_KEY")
	cfg.stripe.secret = os.Getenv("STRIPE_SECRET")


	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	tc := make(map[string]*template.Template)

	app := &application{
		config: cfg,
		infoLog: infoLog,
		errorLog: errorLog,
		templateCache: tc,
		version: VERSION,
	}
}
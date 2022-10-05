package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type config struct {
	addr      string
	staticDir string
}

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	var cfg config
	flag.StringVar(&cfg.addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&cfg.staticDir, "static-dir", "./ui/static", "Path to static assets")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		infoLog:  infoLog,
		errorLog: errorLog,
	}

	srv := &http.Server{
		Addr:     cfg.addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Print("Starting server on ", cfg.addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}

package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP Network Address")
	dsn := flag.String("dsn", "web:pass@tcp(db:3306)/snippetbox?parseTime=true", "Database DSN")

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := sql.Open("mysql", *dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
        log.Fatalf("Error connecting to the database: %v", err)
    }

	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
	}

	
	flag.Parse()

	srv := &http.Server{
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: app.routes(),
	}
	

	infoLog.Printf("starting server on port %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

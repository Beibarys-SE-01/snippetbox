package main

/*import (
	"log"
)

type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
}*/

import (
	"context"
	"flag"
	_ "github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"net/http"
	"os"
	"github.com/Beibarys-SE-01/snippetbox/pkg/models/postgresql"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *postgresql.SnippetModel
}

func main() {

	addr := flag.String("addr", ":4001", "HTTP network address")

	dsn := flag.String("dsn", "postgres://postgres:123@localhost:5432/snippetbox", "PostgreSql data source name")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err1 := openDB(*dsn)
	if err1 != nil {
		errorLog.Fatal(err1)
	}
	defer db.Close()

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		snippets: &postgresql.SnippetModel{DB: db},
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}
	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
func openDB(dsn string) (*pgxpool.Pool, error) {
	db, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

/*func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
	}
	srv := &http.Server{
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: app.routes(), // Call the new app.routes() method
	}
	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}*/
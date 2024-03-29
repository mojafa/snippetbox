package main

import (
	"database/sql"
	"flag"
	"log"
	"mojafa/snippetbox/pkg/models/mysql"
	"net/http"
	"os"

	// However, we need the driver’s init() function to run so that it can register itself with the database/sql package.
	// The trick to getting around this is to alias the package name to the blank identifier. This is standard practice for most of Go’s SQL drivers.

	_ "github.com/go-sql-driver/mysql" // New import
)

// Add a snippets field to the application struct. This will allow us to
// make the SnippetModel object available to our handlers.
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *mysql.SnippetModel
}

func main() {
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name")
	addr := flag.String("addr", ":4000", "HTTP network address")
	// Define a new command-line flag for the MySQL DSN string.
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.LUTC)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.LUTC|log.Lshortfile)

	// To keep the main() function tidy I've put the code for creating a connection // pool into the separate openDB() function below. We pass openDB() the DSN
	// from the command-line flag.
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	// We also defer a call to db.Close(), so that the connection pool is closed
	// before the main() function exits.
	defer db.Close()

	// Initialize a mysql.SnippetModel instance and add it to the application dependencies.
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		snippets: &mysql.SnippetModel{DB: db},
	}

	// Use the http.NewServeMux() function to initialize a new servemux, then
	// register the home function as the handler for the "/" URL pattern.
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	// Because the err variable is now already declared in the code above, we need
	// to use the assignment operator = here, instead of the := 'declare and assign'
	// operator.
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	// sql.Open is used to just initialize the pool
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// verify that everything is set up correctly
	// we use the db.Ping() method to create a connection and
	// check for any errors.

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil

}

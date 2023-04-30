package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	// Define a new command-line flag with the name 'addr', a default value of
	// and some short help text explainig what the flad controls. The value of thhe
	// flag will be stored in the addr variable at the program start-up.
	addr := flag.String("addr", ":4000", "HTTP network address")

	// importantly, we use the flag.Parse() finction to parse the command-line
	//This reads in the command line flag value and assigns it ot the addr
	//variable. You need to call this *before* you use the addr variable
	//otherwise it will always contain the default value of ":4000". If any error is
	//encountered during parsng the application will be terminated.
	flag.Parse()

	// Use log.New() to create a logger for writing information messages. This
	// three parameters: the destination to write the logs to (os.Stdout), a st
	// prefix for message (INFO followed by a tab), and flags to indicate what
	// additional information to include (local date and time). Note that the f
	// are joined using the bitwise OR operator |.

	// f, err := os.OpenFile("/tmp/info.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()

app:= &application{
	errorLog: errorLog,
	infoLog: infoLog,
}


	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// Create a logger for writing error messages in the same way, but use stde
	// the destination and use the log.Lshortfile flag to include the relevant
	// file name and line number.

	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	// Create a file server which serves files out of the "./ui/static" directo
	// Note that the path given to the http.Dir function is relative to the pro
	// directory root.
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	//use the mux.HandleFunc() function to register the file server as the handler
	// for all URL paths that start with "/static/". For example, if the user
	// requests the URL "/static/css/main.css" the file server will search for
	// the "./ui/static/css/main.css" file on the filesystem and serve it if it
	// exists.

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	//Initialize the new http.Server struct. We set the Addr and Handler fields so
	//that the server uses the same network address and routes as before, and the
	//ErrorLog field so that the server now uses the custom errorLog logger in the
	//event of any problems. Note that we also set the ReadTimeout and WriteTimeout
	//fields to 1 second and 10 seconds respectively. This is good practice to
	//prevent slowloris attacks against your application (see Chapter 8 for more
	//details).

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	//The value returned form the flag.String() function is a pointer to the flag
	//string value, not the value itself. So we need to dereference the pointer (prefix it witgh the * symbol) before using it.

	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)

}

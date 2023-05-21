//sql.DB connections pools made up of idle or in-use pools
by default there's no limit on max number of open connections, but the default for the number of idle pool is 2, you can chnage this default using
SetMaxIdleConns
SetMaxOpenConns


db, err := sql.Open("postgres", "postgres://user:pass@localhost:5432/snippetbox sslmode=disable")
if err != nil {
    log.Fatal(err)
}

// Set the maximum number of concurrently open (idle + in-use) connections. Setting this // to less than or equal to 0 will mean there is no maximum limit. If the maximum
// number of open connections is reached and all are in-use when a new connection is
// needed, Go will wait until one of the connections is freed and becomes idle. From a // user perspective, this means their HTTP request will hang until a connection
// is freed.
db.SetMaxOpenConns(100)



CAVEAT: your database itself probably has a hard limit on the maximum number of connections. For example,  the default limit for MySQL is 151

For some applications that behavior might be fine, but in a web application it’s arguably better to immediately log the "too many connections" error message and send a
500 Internal Server Error to the user, rather than having their HTTP request hang and potentially timeout while waiting for a free connection.
That’s why I haven’t used the SetMaxOpenConns() and SetMaxIdleConns() methods in our application, and left the behavior of sql.DB as the default settings.


// Set the maximum number of idle connections in the pool. Setting this
// to less than or equal to 0 will mean that no idle connections are retained.
db.SetMaxIdleConns(5)
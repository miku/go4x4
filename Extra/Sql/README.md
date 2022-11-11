# SQL

## Generic ideas

* Go has low-level support for DB access
* third-party libraries for db implementations
* ORMs exist, but are less central
* lightweight solutions like sqlx seem to be at a sweet spot

## List of drivers

* [https://github.com/golang/go/wiki/SQLDrivers](https://github.com/golang/go/wiki/SQLDrivers)

> To use database/sql you’ll need the package itself, as well as a driver for
> the specific database you want to use.

## Basic Ops

* query (row or rows)
* exec (mutation)

## database/sql

Abstraction layer, limited support for DB specific settings or the underlying
connection pool.

* func (db *DB) SetConnMaxIdleTime(d time.Duration) // maximum amount of time a connection may be idle
* func (db *DB) SetConnMaxLifetime(d time.Duration) // maximum amount of time a connection may be reused
* func (db *DB) SetMaxIdleConns(n int) // maximum number of connections in the idle connection pool
* func (db *DB) SetMaxOpenConns(n int) // maximum number of open connections to the database


### sql.DB

> The sql.DB abstraction is designed to keep you from worrying about how to
> manage concurrent access to the underlying datastore. It’s safe for
> concurrent use by multiple goroutines.

* long-lived object
* lazy (use ping to check connectivity)

Support for transactions, `Tx`, prepared statements. More information:
[https://golang.org/s/sqlwiki](https://golang.org/s/sqlwiki)


### sql.Row and sql.Rows

Return result or a results.

> Row is the result of calling QueryRow to select a single row. 

And:

> Rows is the result of a query. Its cursor starts before the first row of the
> result set. Use Next to advance from row to row. 

### rows.Next

> Next prepares the next result row for reading with the Scan method. 


### Scanner

A single method interface.

> https://pkg.go.dev/database/sql#Scanner

```go
type Scanner interface {
	// Scan assigns a value from a database driver.
	//
	// The src value will be of one of the following types:
	//
	//    int64
	//    float64
	//    bool
	//    []byte
	//    string
	//    time.Time
	//    nil - for NULL values
	//
	// An error should be returned if the value cannot be stored
	// without loss of information.
	//
	// Reference types such as []byte are only valid until the next call to Scan
	// and should not be retained. Their underlying memory is owned by the driver.
	// If retention is necessary, copy their values before the next call to Scan.
	Scan(src any) error
}
```

## sqlx

* third-party, slim layer above standard sql abstraction

Notably, it provides `Get` and `Select` to combine query execution and struct
parsing:
[https://jmoiron.github.io/sqlx/#getAndSelect](https://jmoiron.github.io/sqlx/#getAndSelect)

These time-savers you either `Rows` or `StructScan` under the hood.

> The primary extension on sqlx.Rows is StructScan(), which automatically scans
> results into struct fields. 

```go
    type Place struct {
        Country       string
        City          sql.NullString
        TelephoneCode int `db:"telcode"`
    }
     
    rows, err := db.Queryx("SELECT * FROM place")
    for rows.Next() {
        var p Place
        err = rows.StructScan(&p)
    }
```

Example in todosvc: [https://github.com/miku/todosvc/blob/d1b4a017839ff39f9e140fa76b8094a8c61d4c31/server.go#L45](https://github.com/miku/todosvc/blob/d1b4a017839ff39f9e140fa76b8094a8c61d4c31/server.go#L45).
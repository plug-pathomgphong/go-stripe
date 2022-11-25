package driver

***REMOVED***
	"database/sql"
***REMOVED***

	_ "github.com/go-sql-driver/mysql"
***REMOVED***

func OpenDB(dsn string***REMOVED*** (*sql.DB, error***REMOVED*** {
	db, err := sql.Open("mysql", dsn***REMOVED***
***REMOVED***
		return nil, err
***REMOVED***

	err = db.Ping(***REMOVED***
***REMOVED***
		fmt.Println(err***REMOVED***
		return nil, err
***REMOVED***

	return db, nil
***REMOVED***

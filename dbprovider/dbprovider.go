package dbprovider

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"

	_ "github.com/denisenkom/go-mssqldb"
)

//InitDatabase : generate db
func InitDatabase() *sql.DB {
	log.Println("connecting database.")

	query := url.Values{}
	query.Add("database", "Company")

	u := &url.URL{
		Scheme: "sqlserver",
		User:   url.UserPassword("sa", "1234"),
		Host:   fmt.Sprintf("%s:%d", "localhost", 1433),
		// Path:  instance, // if connecting to an instance instead of a port
		RawQuery: query.Encode(),
	}

	log.Println(u.String())

	condb, err := sql.Open("sqlserver", u.String())
	if err != nil {
		panic(err)
	}
	log.Println("test ping database.")
	if err = condb.Ping(); err != nil {
		panic(err)
	}
	return condb
}

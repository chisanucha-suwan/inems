package utilities

import (
	"database/sql"
	"fmt"
	"net/url"

	_ "github.com/denisenkom/go-mssqldb"
)

func Open(host string, port string, databaseName string, username string, password string) (*sql.DB, error) {
	queryStr := url.Values{}
	queryStr.Add("database", databaseName)

	connectionStr := &url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword(username, password),
		Host:     fmt.Sprintf("%s:%s", host, port),
		RawQuery: queryStr.Encode(),
	}

	return sql.Open("sqlserver", connectionStr.String())
}

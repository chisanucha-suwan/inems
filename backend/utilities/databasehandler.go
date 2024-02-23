package utilities

import (
	"database/sql"
	"fmt"
	"net/url"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/doug-martin/goqu/v9"
)

// init initializes the default prepared statement setting for the Goqu package.
func init() {
	// Performance: Reuses optimized execution plans for faster queries.
	// Security: Prevents SQL injection by separating parameters from the query.
	// Readability: Improves code clarity by showing parameter usage explicitly.
	goqu.SetDefaultPrepared(true)
}

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

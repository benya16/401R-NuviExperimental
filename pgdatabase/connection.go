package pgdatabase

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
)

type Connection struct {
	database* sql.DB
}

func (c *Connection) Connect() {
	db, err := sql.Open("postgres", "postgres://go:gogo2017@localhost/nuvisocialthreat?sslmode=disable")
	sqlError(err, "Error at Connect()")
	c.database = db
}

func (c *Connection) IsConnected() bool {
	if c.database == nil {
		return false
	}
	return true
}

func (c *Connection) Query(sqlStatement string) *sql.Rows{
	rows, err := c.database.Query(sqlStatement)
	sqlError(err, "Error at Query()")

	return rows
}

func (c *Connection) Execute(sqlStatement string){
	_, err := c.database.Exec(sqlStatement)
	sqlError(err, "Error at Execute")
}

func (c *Connection) Close() {
	c.database.Close()
	c.database = nil
}

func (c *Connection) Prepare(query string) *sql.Stmt {
	result, err := c.database.Prepare(query)
	sqlError(err, "Error at Prepare()")

	return result
}

func (c *Connection) Test() {
	if c.database == nil {
		fmt.Println("database is nil")
	} else {
		fmt.Println("not nil")
	}
}
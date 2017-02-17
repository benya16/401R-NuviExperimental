package Database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Connection struct {
	database* sql.DB
}

func (c *Connection) Connect() {
	db, err := sql.Open("postgres", "postgres://go:gogo2017@localhost/nuvisocialthreat?sslmode=disable")
	sqlError(err)
	c.database = db
}

func (c *Connection) Query(sqlStatement string) *sql.Rows{
	rows, err := c.database.Query(sqlStatement)
	sqlError(err)

	return rows
}

func (c *Connection) Execute(sqlStatement string){
	_, err := c.database.Exec(sqlStatement)
	sqlError(err)
}

func (c *Connection) Close() {
	c.database.Close()
}

func (c *Connection) Prepare(query string) *sql.Stmt {
	result, err := c.database.Prepare(query)
	sqlError(err)

	return result
}
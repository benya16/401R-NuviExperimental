package Database

import "fmt"

func CreateTables() {
	conn := new(Connection)
	createStatement := readSQLFile("resources/sql/createTables.sql")
	fmt.Println(createStatement)
	conn.Connect()
	conn.Execute(createStatement)
	conn.Close()
}
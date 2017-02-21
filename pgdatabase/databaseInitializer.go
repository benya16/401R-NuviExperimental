package pgdatabase

func CreateTables() {
	conn := new(Connection)
	createStatement := readSQLFile("resources/sql/createTables.sql")
	conn.Connect()
	conn.Execute(createStatement)
	conn.Close()
}
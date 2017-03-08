package pgdatabase

import "database/sql"

type DAO struct {
	connection *Connection
}

func NewDAO() *DAO {
	dao := new(DAO)
	dao.connection = new(Connection)

	return dao
}

func (this *DAO) AddPost(post []byte) {
	var transaction bool = false
	if this.connection.IsConnected() {
		transaction = true
	}
	if !transaction {
		this.connection.Connect()
	}
	stmt := this.connection.Prepare(readSQLFile("resources/sql/insertPost.sql"))
	stmt.Exec(post)
	if !transaction {
		this.connection.Close()
	}
}

func (this *DAO) GetGeoHash(geohash string) *sql.Rows{
	var transaction bool
	if this.connection.IsConnected() {
		transaction = true
	}
	if !transaction {
		this.connection.Connect()
	}
	stmt := this.connection.Prepare(readSQLFile("resources/sql/getGeoHash.sql"))
	result, err := stmt.Query(geohash)
	sqlError(err, "Error in GetGeoHash")

	return result
}

func (this *DAO) Test() {
	this.connection.Test()
}
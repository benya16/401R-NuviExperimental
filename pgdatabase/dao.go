package pgdatabase

type DAO struct {
	connection *Connection
}

func NewDAO() *DAO {
	dao := new(DAO)
	dao.connection = new(Connection)

	return dao
}

func (d *DAO) AddPost(post string) {
	d.connection.Connect()
	stmt := d.connection.Prepare(readSQLFile("resources/sql/insertPost.sql"))
	stmt.Exec("geohash", []byte(post))
	d.connection.Close()
}


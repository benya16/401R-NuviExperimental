package main

import (
	"../pgdatabase"
)

func main() {
	//pgdatabase.CreateTables()
	dao := pgdatabase.NewDAO()
	dao.AddPost("this is a test post")
}

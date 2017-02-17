package main

import (
	"../Database"
)

func main() {
	//Database.CreateTables()
	dao := Database.NewDAO()
	dao.AddPost("this is a test post")
}

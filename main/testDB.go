package main

import (
	"../pgdatabase"
	"fmt"
	"time"
)

func main() {
	//pgdatabase.CreateTables()
	dao := pgdatabase.NewDAO()
	//dao.AddPost("this is a test post")
	//dao.Test()
	result := dao.GetGeoHash("geohash")
	fmt.Println(result.Columns())
	var (
		uuid string
		collected time.Time
		geo_hash string
		post []byte
	)
	for result.Next() {
		result.Scan(&uuid, &collected, &geo_hash, &post)
		fmt.Println(uuid, collected, geo_hash, string(post))
	}
}

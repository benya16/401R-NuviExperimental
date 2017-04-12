package main

import (
	"database/sql"
	"fmt"
	"os"
	"../pgdatabase"
	"../filter"
	"../models"
	"encoding/json"
)

func main() {
	//nuvi, err := sql.Open("postgres", "postgres://go:gogo2017@localhost/nuvisocialthreat?sslmode=disable")
	//sqlError(err, "Error connecting to socialthreat")
	dao := pgdatabase.NewDAO()
	twitterFilter := new(filter.Filter)
	twitterFilter.InitFilter("danger.csv")
	manual, err := sql.Open("postgres", "postgres://go:gogo2017@localhost/manualthreats?sslmode=disable")
	sqlError(err, "Error connecting to socialthreat")

	rows, err := manual.Query("select uuid, threat, post from post")
	var id string
	var threat bool
	var storedPost []byte
	var post models.Post
	for rows.Next() {
		err := rows.Scan(&id, &threat, &storedPost)
		sqlError(err, "Error Scanning result set")

		json.Unmarshal(storedPost, &post)
		dao.AddRawPost(id, storedPost)
		processed := twitterFilter.Preprocess(&post)
		dao.AddProcessedPost(id, processed)
		if threat {
			dao.LabelThreat(id, true)
		} else {
			dao.LabelThreat(id, false)
		}
	}
}

func sqlError(err error, message string) {
	if err != nil {
		fmt.Println("SQL Error: ", message, "-> ", err.Error())
		os.Exit(1)
	}
}
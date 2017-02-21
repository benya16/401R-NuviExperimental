package pgdatabase

import (
	"io/ioutil"
	"fmt"
	"os"
)

func readSQLFile(fileName string) string {
	data, err := ioutil.ReadFile(fileName)
	sqlError(err, "Error while reading sql file")

	return string(data)
}

func sqlError(err error, message string) {
	if err != nil {
		fmt.Println("SQL Error: ", message, "-> ", err.Error())
		os.Exit(1)
	}
}
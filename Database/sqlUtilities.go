package Database

import (
	"io/ioutil"
	"fmt"
	"os"
)

func readSQLFile(fileName string) string {
	data, err := ioutil.ReadFile(fileName)
	sqlError(err)

	return string(data)
}

func sqlError(err error){
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
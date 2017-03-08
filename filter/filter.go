package filter

import (
	"io/ioutil"
	"strings"
	"regexp"
	"../models"
	"fmt"
)

var dangerDictionary []string
var regex *regexp.Regexp
var exceptionRegex *regexp.Regexp

type Filter struct {
	exceptionFilter bool
}

func (d* Filter) InitFilter(dictionaryFileName string) {
	d.exceptionFilter = false
	data, _ := ioutil.ReadFile(dictionaryFileName)
	rows := strings.Split(string(data), ",")
	for _, row := range rows {
		entries := strings.Split(row, ",")
		for _, entry := range entries {
			if entry != "" {
				dangerDictionary = append(dangerDictionary, entry)
			}
		}
	}
	regexString := "(?i)"
	for i := 0; i < len(dangerDictionary); i++ {
		if i > 0 {
			regexString += "|\\s+#?"
		}
		regexString += strings.TrimSpace(dangerDictionary[i]) + "\\s+"
	}
	//fmt.Println(dangerDictionary)
	//fmt.Println(regexString)
	regex = regexp.MustCompile(regexString)
}

func (d* Filter) SetExceptionsFilter(dictionaryFileName string) {
	d.exceptionFilter = true
	var exceptionDictionary []string
	data, _ := ioutil.ReadFile(dictionaryFileName)
	rows := strings.Fields(string(data))
	for _, row := range rows {
		entries := strings.Split(row, ",")
		for _, entry := range entries {
			if entry != "" {
				exceptionDictionary = append(exceptionDictionary, entry)
			}
		}
	}
	//exceptionDictionary := strings.Split(string(data), ",")
	regexString := "(?i)"
	//fmt.Println(exceptionDictionary)
	for i := 0; i < len(exceptionDictionary); i++ {
		if i > 0 {
			regexString += "|\\s+"
		}
		regexString += strings.TrimSpace(exceptionDictionary[i]) + "\\s+"
	}
	//fmt.Println(regexString)
	exceptionRegex = regexp.MustCompile(regexString)
}

func (d* Filter) ContainsDangerWord(post models.Post) bool {
	if post.Language == "en" {
		slice := regex.FindAllString(post.Raw_body_text, -1)
		if len(slice) > 0 {
			fmt.Print("Matched on: ")
			fmt.Println(slice)
			if d.exceptionFilter {
				return !exceptionRegex.Match([]byte(post.Raw_body_text))
			} else {
				return true
			}
		}
	}

	return false
}

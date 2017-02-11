package filter

import (
	"io/ioutil"
	"strings"
	"regexp"
)

var dangerDictionary []string
var regex *regexp.Regexp

type Filter struct {}

func (d* Filter) InitFilter(dictionaryFileName string) {
	data, _ := ioutil.ReadFile(dictionaryFileName)
	dangerDictionary = strings.Fields(string(data))
	regexString := "(?i)"
	for i := 0; i < len(dangerDictionary); i++ {
		if i > 0 {
			regexString += "|"
		}
		regexString += dangerDictionary[i]
	}
	regex = regexp.MustCompile(regexString)
}

func (d* Filter) ContainsDangerWord(bodyText string) bool {
	return regex.Match([]byte(bodyText))
}

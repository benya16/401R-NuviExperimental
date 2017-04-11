package filter

import (
	"io/ioutil"
	"strings"
	//"regexp"
	"regexp"
)

var dangerDictionary []string

var regex *regexp.Regexp

var Dictionary *Trie

type Filter struct {}

func (d* Filter) InitFilter(dictionaryFileName string) {
	Dictionary = NewTrie()
	data, _ := ioutil.ReadFile(dictionaryFileName)
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		phrases := strings.Split(line, ",")
		for _, phrase := range phrases {
			phrase = strings.TrimSpace(phrase)
			if phrase != "" {
				dangerDictionary = append(dangerDictionary, phrase)
			}
		}
	}
	for i := 0; i < len(dangerDictionary); i++ {
		Dictionary.AddWordWIthDerivation(dangerDictionary[i], false)
	}

	//regexString := "(?i)"
	//for i := 0; i < len(dangerDictionary); i++ {
	//	if i > 0 {
	//		regexString += "|"
	//	}
	//	regexString += dangerDictionary[i]
	//}
	//regex = regexp.MustCompile(regexString)

}

func (d* Filter) InitExceptions(dictionaryFileName string) {
	data, _ := ioutil.ReadFile(dictionaryFileName)
	dangerDictionary = strings.Fields(string(data))

	for i := 0; i < len(dangerDictionary); i++ {
		Dictionary.AddWordWIthDerivation(dangerDictionary[i], true)
	}
}

func (d* Filter) ContainsDangerWord(bodyText string) bool {
	//result := regex.FindAllString(bodyText, -1)
	//if len(result) > 0 {
	//	fmt.Println(result)
	//	return true
	//}
	//return false
	returnVar := Dictionary.isDangerousSentance(bodyText)
	if(returnVar.size == 0) {
		return false
	}

	return true
}

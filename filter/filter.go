package filter

import (
	"io/ioutil"
	"strings"
	//"regexp"
)

var dangerDictionary []string



//var regex *regexp.Regexp

var trie *Trie

type Filter struct {}




func (d* Filter) InitFilter(dictionaryFileName string) {


	trie = NewTrie()


	data, _ := ioutil.ReadFile(dictionaryFileName)
	dangerDictionary = strings.Fields(string(data))



	for i := 0; i < len(dangerDictionary); i++ {
		trie.AddWordWIthDerivation(dangerDictionary[i], false)
	}


/*
	regexString := "(?i)"
	for i := 0; i < len(dangerDictionary); i++ {
		if i > 0 {
			regexString += "|"
		}
		regexString += dangerDictionary[i]
	}
	regex = regexp.MustCompile(regexString)
	*/
}



func (d* Filter) InitExceptions(dictionaryFileName string) {

	data, _ := ioutil.ReadFile(dictionaryFileName)
	dangerDictionary = strings.Fields(string(data))


	for i := 0; i < len(dangerDictionary); i++ {
		trie.AddWordWIthDerivation(dangerDictionary[i], true)
	}
}




func (d* Filter) ContainsDangerWord(bodyText string) bool {


	returnVar := trie.isDangerousSentance(bodyText)
	if(returnVar.size == 0) {
		return false
	}
	return true


}

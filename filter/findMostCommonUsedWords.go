package filter

import (
	"regexp"
	"strings"
)






func addWords(post string) [100]node {



	var TopWords [100]node

	for i := 0; i < 100; i++ {
		TopWords[i] = node{word: "", count: -1}
	}


	trie := NewNumericTrie()

	//sanitize first
	re := regexp.MustCompile("/.|,|!|/?")



	words := strings.Fields(re.ReplaceAllString(post, " "))

	for _, word :=range words {
		var number = trie.addWord(word)

		var lowestIndex = -1
		for index, point := range TopWords {
			if(point.count == -1) {
				lowestIndex = index
				break
			}

			if(lowestIndex == -1 || TopWords[lowestIndex].count < point.count) {
				lowestIndex = index
			}
		}

		if(number > TopWords[lowestIndex].count) {
			TopWords[lowestIndex].count = number
			TopWords[lowestIndex].word = word
		}
	}

	return TopWords


}

type node struct {
	word string
	count int
}

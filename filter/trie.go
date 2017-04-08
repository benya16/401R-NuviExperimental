package filter

type Link struct {
	value rune
	link *Trie
}

type Trie struct {
	childNodes []Link
	stringEnd bool
	word string
	exception bool

}





func findRuneLink(links []Link, value rune) (*Trie, bool) {		//this is the original way we sorted through tries.  Does not work for multiple words


	for _, link := range links {
		if link.value == value {
			return link.link, true
		}
	}
	return nil, false
}





func (r *Trie) findRuneInSentance(links []Link, value rune) (*Trie, bool) {		//this is try 2, works for sentances

	//first look to see if the symbol can be processed.  spaces can be consumed during this step if it is part of the word
	for _, link := range links {
		if link.value == value {
			return link.link, true
		}
	}

	//otherwise, if there is a space go back to root
	if value == ' ' {
		return r, true
	}


	//otherwise, if there is nowhere else we can go to, then say we can't go anywhere
	return nil, false

}




//this was the default way to add a word.
func (r *Trie) ExistsOrAdd(s string) bool {
	check := true
	i := r
	for _, runeValue := range s {
		ti, ok := findRuneLink(i.childNodes,runeValue)
		if !ok {
			ti = new(Trie)
			ti.childNodes = make([]Link, 0)
			i.childNodes = append(i.childNodes,Link{value: runeValue, link: ti})
		}
		i = ti
	}
	if !i.stringEnd {
		i.stringEnd = true
		check = false
	}
	return check
}
	//a container for storing the words that we found to be dangerous
func NewWordSet() *WordSet {
	return &WordSet{set: make(map[int]string), size: 0}
}

type WordSet struct {
	set map[int]string
	size int
}

func (words *WordSet) Add(s string) bool {
	words.set[words.size] = s
	words.size = words.size + 1
	return true
}

func (words *WordSet) Get(i int) string {
	return words.set[i]
}





//USE THIS FUNCTION WHEN USING THE TRIE TO FIND DANGEROUS WORDS
func (r *Trie) AddWordWIthDerivation(s string, exception bool) bool {


	//add in generator to add in all of the words with punctuation

	if r.addWord(s, exception) {
		r.addWord(s + "!", exception)
		r.addWord(s + ".", exception)
		r.addWord(s + ",", exception)
		r.addWord(s + "?", exception)
		return true
	}

	return false


}













//returns true if the word was added, false if word was already added
func (r *Trie) addWord(s string, exception bool) bool {
	i := r
	for _, symbolValue := range s {
		ti, ok := findRuneLink(i.childNodes,symbolValue)
		if !ok {
			ti = new(Trie)
			ti.childNodes = make([]Link, 0)
			i.childNodes = append(i.childNodes, Link{value: symbolValue, link: ti})
		}
		i = ti
	}
	if !i.stringEnd {
		i.stringEnd = true
		i.exception = exception
		i.word = s
		return true
	}
	return false
}






func NewTrie() *Trie {
	return &Trie{stringEnd: true, childNodes: make([]Link, 0)}
}


//can only take one word at a time
func (r *Trie) isDangerous(s string) bool  {
	i := r
	for _, symbolValue := range s {
		ti, ok := findRuneLink(i.childNodes, symbolValue)
		if ok {
			i = ti
		} else {
			return false
		}
	}
	if i.stringEnd {
		if i.exception {
			return false
		}
		return true
	}
	return false;
}


//takes a whole sentance.
func (r *Trie) isDangerousSentance(sentance string) WordSet {

	var s = sentance + " "

	if(len(s) > 25000) {
		return *NewWordSet()
	}



	set := NewWordSet()


	var foundWordEnd = false
	var foundWord = ""

	i := r
	for _, symbolValue := range s {
		ti, ok := findRuneLink(i.childNodes, symbolValue)
		if ok {
			i = ti
			foundWordEnd = false
			foundWord = ""
			if(ti.stringEnd) {
				if(ti.exception) {
					return *NewWordSet()
				}
				foundWordEnd = true
				foundWord = ti.word

			}
		} else {
			//check for terminal character
			if(foundWordEnd && symbolValue == ' ') {
				set.Add(foundWord)
			}
			i = r
			foundWordEnd = false
			foundWord = ""
		}
	}


	return *set
}






type TrieNumeric struct {
	childNodes []NumericLink
	repetitionsOfWord int
}


type NumericLink struct {
	value rune
	link *TrieNumeric
}



func findNumericRuneLink(links []NumericLink, value rune) (*TrieNumeric, bool) {


	for _, link := range links {
		if link.value == value {
			return link.link, true
		}
	}
	return nil, false
}


//PUT IN SINGLE, SANITIZED WORDS (no punctuation or spaces)
func (r *TrieNumeric) addWord(s string) int {
	//the problem with this is that we need to handle imperfect words, EG
	//with punctuation, spaces, and so forth
	//we could either sanitize it, or...
	//we should sanitize it


	i := r
	for _, symbolValue := range s {
		ti, ok := findNumericRuneLink(i.childNodes,symbolValue)
		if !ok {
			ti = new(TrieNumeric)
			ti.childNodes = make([]NumericLink, 0)
			i.childNodes = append(i.childNodes, NumericLink{value: symbolValue, link: ti})
		}
		i = ti
	}
	i.repetitionsOfWord = i.repetitionsOfWord + 1
	return i.repetitionsOfWord
}














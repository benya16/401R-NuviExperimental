package filter

import (
	"io/ioutil"
	"strings"
	//"regexp"
	"regexp"
	"../models"
	"fmt"
)

var dangerDictionary []string

//var regex *regexp.Regexp

//var Dictionary *Trie

type Filter struct {
	dictionary *Trie
	regex *regexp.Regexp
}

func (d* Filter) InitFilter(dictionaryFileName string) {
	d.dictionary = NewTrie()
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
		d.dictionary.AddWordWIthDerivation(dangerDictionary[i], false)
	}

	regexString := "(?i)"
	for i := 0; i < len(dangerDictionary); i++ {
		if i > 0 {
			regexString += "|"
		}
		regexString += dangerDictionary[i]
	}
	d.regex = regexp.MustCompile(regexString)

}

func (d* Filter) InitExceptions(dictionaryFileName string) {
	data, _ := ioutil.ReadFile(dictionaryFileName)
	dangerDictionary = strings.Fields(string(data))

	for i := 0; i < len(dangerDictionary); i++ {
		d.dictionary.AddWordWIthDerivation(dangerDictionary[i], true)
	}
}

func (d* Filter) ContainsDangerWord(bodyText string) bool {
	result := d.regex.FindAllString(bodyText, -1)
	if len(result) > 0 {
		fmt.Println(result)
		return true
	}
	return false
	//returnVar := d.dictionary.isDangerousSentance(bodyText)
	//if(returnVar.size == 0) {
	//	return false
	//}
	//
	//return true
}

//func (d* Filter) GetMatchedWords(bodyText string) []string {
//	matches := regex.FindAllString(bodyText, -1)
//	sort.Strings(matches)
//	for i, match := range matches {
//		matches[i] = strings.ToLower(match)
//	}
//	return matches
//}

func (d* Filter) Preprocess(post *models.Post) *models.ProcessedPost {
	processed := new(models.ProcessedPost)
	processed.PostLength = uint(len(post.Raw_body_text))
	processed.LikeCount = post.Like_count
	processed.FollowersCount = post.Author_followers_count
	processed.FriendCount = post.Author_friends_count
	processed.HashtagCount = uint(len(post.Hashtags))
	processed.RetweetCount = post.Retweet_count
	processed.IsRetweet = post.Is_reshare
	processed.KloutScore = post.Author_klout_score
	exclamationCount := 0
	for _, c := range post.Raw_body_text {
		character := string(c)
		if character == "!" {
			exclamationCount++
		}
	}
	processed.ExclaimationCount = uint(exclamationCount)

	//take the raw body text, run it through the filter


	dangerWords := d.dictionary.isDangerousSentance(post.Raw_body_text)

	processed.Shooter = dangerWords.Contains("shooter")
	processed.ActiveShooter = dangerWords.Contains("active shooter")
	processed.Attack = dangerWords.Contains("attack")
	processed.Bomb = dangerWords.Contains("bomb")
	processed.BombThreat = dangerWords.Contains("bomb threat")
	processed.Breaking = dangerWords.Contains("breaking")
	processed.Danger = dangerWords.Contains("danger")
	processed.Dead = dangerWords.Contains("dead")
	processed.Gunman = dangerWords.Contains("gunman")
	processed.Killing = dangerWords.Contains("killing")
	processed.Rape = dangerWords.Contains("rape")
	processed.Shooting = dangerWords.Contains("shooting")
	processed.Stabbing = dangerWords.Contains("stabbing")
	processed.Terrorist = dangerWords.Contains("terrorist")
	processed.Warning = dangerWords.Contains("warning")

	return processed
}
package main

import (
	"io/ioutil"
	"os"
	"strings"
	"math/rand"
	"strconv"
	"time"
	"encoding/json"
)

var SafeDictionary []string
var DangerDictionary []string
var NumToGenerate int
var PercentDangerous float64
var JsonFileName string

type Post struct {
	Activity_url string
	Author_favorites_count uint
	Author_followers_count uint
	Author_friends_count uint
	Author_klout_score uint
	Author_picture_url string
	Author_posts_count uint
	Author_profile_url string
	Author_real_name string
	Author_username string
	Bio string
	Cleaned_body_text string
	Country string
	Country_code string
	Embedded_urls []string
	Hashtags []string
	Historical_search  bool
	Is_reshare bool
	Language string
	Latitude float64
	Like_count uint
	Location_display_name string
	Longitude float64
	Mentions  []string
	Meta_data []interface{}
	Network string
	Normalized_urls []string
	Parent struct{
		Parent_author   string
		Parent_author_profile_url string
		Parent_author_reach int
		Parent_body_text string
		Parent_created_at string
		Parent_social_source_id string
	}
	Post_created_at string
	Post_media  struct {
		Media_url string
		Url string
		Display_url string
		Media_type string
		Video_url string
	}
	Raw_body_text string
	Region string
	Region_code string
	Retweet_count uint
	Social_monitor_sources []struct {
		Company_uid string
		Monitor_uid string
		Keywords map[string]uint
	}
	Social_monitor_uids []string
	Social_source_uid string
	Source string
	Topic_monitor_id string
}

func main() {
	rand.Seed(time.Now().Unix())

	data, _ := ioutil.ReadFile(os.Args[1])
	SafeDictionary = strings.Fields(string(data))
	data, _ = ioutil.ReadFile(os.Args[2])
	DangerDictionary = strings.Fields(string(data))
	num, _ := strconv.Atoi(os.Args[3])
	NumToGenerate = int(num)
	dumb, _ := strconv.ParseFloat(os.Args[4], 64)
	PercentDangerous = float64(dumb)
	JsonFileName = string(os.Args[5])

	var posts []Post

	for i := 0; i < NumToGenerate; i++ {
		post := new(Post)
		if float64(i) < float64(NumToGenerate) * PercentDangerous {
			post.Raw_body_text = getDangerSentence()
		} else {
			post.Raw_body_text = getSafeSentence()
		}
		posts = append(posts, *post)
	}

	data, _ = json.Marshal(posts)
	ioutil.WriteFile(JsonFileName, data, 0644)
}

func getSafeSentence() string {
	var result string
	numWords := rand.Intn(35) + 15
	for j := 0; j < numWords; j++ {
		result += SafeDictionary[rand.Intn(len(SafeDictionary))] + " "
	}
	return result
}

func getDangerSentence() string {
	var result string
	var words []string
	var shuffledWords []string
	numWords := rand.Intn(33) + 14
	numDangerWords := rand.Intn(2) + 1
	for i := 0; i < numDangerWords; i++ {
		words = append(words, DangerDictionary[rand.Intn(len(DangerDictionary))])
	}
	for j := 0; j < numWords; j++ {
		words = append(words, SafeDictionary[rand.Intn(len(DangerDictionary))])
	}
	perm := rand.Perm(len(words))
	for _, i := range perm {
		shuffledWords = append(shuffledWords, words[i])
	}
	for i := 0; i < len(shuffledWords); i++ {
		result += shuffledWords[i] + " "
	}
	return result
}
package main

import (
	"io/ioutil"
	"os"
	"strings"
	"math/rand"
	"strconv"
	"time"
	"encoding/json"
	"./models"
)

var SafeDictionary []string
var DangerDictionary []string
var NumToGenerate int
var PercentDangerous float64
var JsonFileName string

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

	var posts []models.Post

	for i := 0; i < NumToGenerate; i++ {
		post := new(models.Post)
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
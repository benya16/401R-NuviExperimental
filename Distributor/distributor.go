package distributor

import (
	"io/ioutil"
	"fmt"
	"os"
	"encoding/json"
)

type Post struct {
	Url    	  	string    `json:"activity_url"`
	Text   	  	string    `json:"cleaned_body_text"`
	RawText   	string    `json:"raw_body_text"`
	Country	  	string    `json:"country"`
	CountryCode	string    `json:"country_code"`
	Region		string    `json:"region"`
	RegionCode	string    `json:"region_code"`
	Language	string    `json:"language"`
	Latitude	float64   `json:"latitude"`
	Longitude	float64   `json:"longitude"`
}

func (p Post) toString() string {
	return toJson(p)
}

func toJson(p interface{}) string {
	bytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return string(bytes)
}

type Distributor struct {
	Posts []Post
}

func getPosts(d Distributor)  {
	raw, err := ioutil.ReadFile("data/normalized-payloads.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c []Post
	json.Unmarshal(raw, &c)
	d.Posts = c
}
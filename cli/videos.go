package main

import (
	"encoding/json"
	"io/ioutil"
)

type video struct {
	Id string
	Title string
	Description string
	ImageUrl string
	Url string
}

// Function that reads a local json file (ioutil.REeadFile) containing videos that 
// use the above data structure. Videos are read top a var called fileBytes 
// (+ subsequent err handling). The json module is then used to Unmarshal the fileBytes, 
// pointing it to a return variable called videos (slice of videos). This second process also involves err 
// handling via the use of panic(err).
func getVideos()(videos []video) {
	fileBytes, err := ioutil.ReadFile("./videos.json")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(fileBytes, &videos)

	if err != nil {
		panic(err)
	}

	return videos
}

// Takes in a slice of videos
// The videos slice is marshalled (converted) to a byte array using json.Marshall (+ err handling)
// Write the byte array back to the json file
func saveVideos(videos []video)() {
	videoBytes, err := json.Marshal(videos)

	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("./videos.json", videoBytes, 0644)

	if err != nil {
		panic(err)
	}
}


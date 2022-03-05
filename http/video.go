package main

import (
	"encoding/json"
	"io/ioutil"
)

// type of video containing the following fields: id, title, description, url, imageurl
type video struct {
	Id          string
	Title       string
	Description string
	Url         string
	Imageurl    string
}

// 1. fucntion called getVideos that takes nothing and returns a slice of videos
// 2. Before you can return a slice of objects of type video, you'll need to use the ioutil package to read in bytes from a json file
// 3. Now that you have a slice of bytes, you'll need to use json.Unmarshal to convert the bytes and store it in your videos array
func getVideos() (videos []video) {
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

// 1. Function called saveVideos that takes in a slice of videoes and returns nothing
// 2. Marshall the videos input into an array of bytes
// 3. Write the videoBytes to the videos.json file using ioutil.WriteFile, making sure to assign it permissions of 0644 if it doesn't exist
func saveVideos(videos []video) {
	videoBytes, err := json.Marshal(videos)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("./videos.json", videoBytes, 0644)
	if err != nil {
		panic(err)
	}
}
